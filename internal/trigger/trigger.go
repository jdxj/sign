package trigger

import (
	"fmt"
	"sync"

	"github.com/go-mysql-org/go-mysql/canal"
	"github.com/panjf2000/ants/v2"
	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/pkg/mq"
	"github.com/jdxj/sign/internal/trigger/dao/specification"
)

func New(conf config.DB) *Trigger {
	trg := &Trigger{
		cron:    cron.New(),
		wg:      &sync.WaitGroup{},
		specIDs: make(map[int64]struct{}),
	}
	// Goroutine Pool
	pool, err := ants.NewPool(1000, ants.WithNonblocking(true))
	if err != nil {
		panic(err)
	}
	trg.gPool = pool

	// Canal
	canalCfg := canal.NewDefaultConfig()
	canalCfg.Addr = fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	canalCfg.User = conf.User
	canalCfg.Password = conf.Pass
	canalCfg.Dump.TableDB = conf.Dbname
	canalCfg.Dump.Tables = []string{"specification"}
	canalCfg.Dump.ExecutionPath = ""
	can, err := canal.NewCanal(canalCfg)
	if err != nil {
		panic(err)
	}
	can.SetEventHandler(trg)
	trg.canal = can

	tq, err := mq.NewTaskQueue()
	if err != nil {
		panic(err)
	}
	trg.tq = tq

	return trg
}

type Trigger struct {
	cron  *cron.Cron
	gPool *ants.Pool

	canal *canal.Canal
	canal.DummyEventHandler

	wg *sync.WaitGroup
	tq *mq.TaskQueue

	specIDs map[int64]struct{}
}

func (trg *Trigger) Start() {
	trg.cron.Start()
	// 全量恢复
	err := trg.recoverJob()
	if err != nil {
		panic(err)
	}
	// 增量添加
	trg.syncJob()
}

func (trg *Trigger) Stop() {
	// 注意 stop 的顺序
	trg.canal.Close()
	<-trg.cron.Stop().Done()
	trg.gPool.Release()
	trg.tq.Stop()

	trg.wg.Wait()
}

func (trg *Trigger) addJob(spec string) error {
	j := &job{
		spec:  spec,
		gPool: trg.gPool,
		tq:    trg.tq,
	}
	_, err := trg.cron.AddJob(spec, j)
	return err
}

func (trg *Trigger) recoverJob() error {
	specs, err := specification.Find(nil)
	if err != nil {
		return err
	}
	for _, spec := range specs {
		err = trg.addJob(spec.Spec)
		if err != nil {
			logger.Errorf("add job failed: %s", err)
		} else {
			trg.specIDs[spec.SpecID] = struct{}{}
			logger.Debugf("add job success: %s", spec.Spec)
		}
	}
	return nil
}

func (trg *Trigger) OnRow(e *canal.RowsEvent) error {
	if e.Table.Name != "specification" {
		return nil
	}

	logger.Debugf("%s %v", e.Action, e.Rows)
	switch e.Action {
	default:
		return nil
	case canal.InsertAction:
	}

	if len(e.Rows) < 1 {
		return fmt.Errorf("invalid RowsEvent")
	}
	sp, err := parseSpec(e.Rows[0])
	if err != nil {
		logger.Errorf("parse spec failed: %s", err)
		return nil
	}
	// 避免旧的 binlog
	_, ok := trg.specIDs[sp.SpecID]
	if ok {
		return nil
	}

	err = trg.addJob(sp.Spec)
	if err != nil {
		logger.Errorf("add job failed, err: %s, specID: %d, spec: %s",
			err, sp.SpecID, sp.Spec)
		return nil
	}
	trg.specIDs[sp.SpecID] = struct{}{}
	return nil
}

func parseSpec(columns []interface{}) (*specification.Specification, error) {
	if len(columns) != 2 {
		return nil, fmt.Errorf("specification table may be changed")
	}
	specID, ok := columns[0].(int64)
	if !ok {
		return nil, fmt.Errorf("invalid specID: %v", columns[0])
	}
	spec, ok := columns[1].(string)
	if !ok {
		return nil, fmt.Errorf("invalid spec: %v", columns[1])
	}
	sp := &specification.Specification{
		SpecID: specID,
		Spec:   spec,
	}
	return sp, nil
}

func (trg *Trigger) syncJob() {
	trg.wg.Add(1)
	go func() {
		defer trg.wg.Done()

		err := trg.canal.Run()
		if err != nil {
			logger.Errorf("canal run err: %s", err)
		}
	}()
}

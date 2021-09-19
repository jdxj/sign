package trigger

import (
	"github.com/panjf2000/ants/v2"
	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/trigger/dao/specification"
)

func New() *Trigger {
	trg := &Trigger{
		cron: cron.New(),
	}
	pool, err := ants.NewPool(1000, ants.WithNonblocking(true))
	if err != nil {
		panic(err)
	}
	trg.gPool = pool
	return trg
}

type Trigger struct {
	cron  *cron.Cron
	gPool *ants.Pool
}

func (trg *Trigger) Start() {
	trg.cron.Start()
	err := trg.RecoverJob()
	if err != nil {
		panic(err)
	}
}

func (trg *Trigger) Stop() {
	<-trg.cron.Stop().Done()
}

func (trg *Trigger) addJob(spec string) error {
	j := &job{
		spec:  spec,
		gPool: trg.gPool,
	}
	_, err := trg.cron.AddJob(spec, j)
	return err
}

func (trg *Trigger) RecoverJob() error {
	specs, err := specification.Find(nil)
	if err != nil {
		return err
	}
	for _, spec := range specs {
		err = trg.addJob(spec.Spec)
		if err != nil {
			logger.Errorf("add job failed: %s", err)
		} else {
			logger.Debugf("add job success: %s", spec.Spec)
		}
	}
	return nil
}

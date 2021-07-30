package task

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/pkg/bot"
	"github.com/jdxj/sign/internal/task/bili"
	"github.com/jdxj/sign/internal/task/common"
	"github.com/jdxj/sign/internal/task/hpi"
	"github.com/jdxj/sign/internal/task/stg"
)

var (
	ic      = cron.New()
	manager = NewManager()
)

func Add(t *Task) error {
	return manager.Add(t)
}

func Run() {
	addFunc(ic)
	ic.Run()
}

func Start() {
	addFunc(ic)
	ic.Start()
}

func addFunc(c *cron.Cron) {
	_, _ = c.AddFunc("0 20 * * *", manager.Run)
}

type Task struct {
	ID     string
	Domain int
	Types  []int
	Key    string
	Client *http.Client
}

func NewManager() *Manager {
	return &Manager{
		tasks: make(map[*http.Client]*Task),
	}
}

type Manager struct {
	tasks map[*http.Client]*Task
}

func verifyTypes(types []int) error {
	for _, t := range types {
		_, ok := common.TypeMap[t]
		if !ok {
			return fmt.Errorf("%w: %d", common.ErrorUnsupportedType, t)
		}
	}
	return nil
}

func (m *Manager) Add(t *Task) (err error) {
	switch t.Domain {
	case common.BiliDomain:
		t.Client, err = bili.Auth(t.Key)

	case common.HPIDomain:
		t.Client, err = hpi.Auth(t.Key)

	case common.STGDomain:
		t.Client, err = stg.Auth(t.Key)

	default:
		err = common.ErrorUnsupportedDomain
	}
	if err != nil {
		return
	}

	err = verifyTypes(t.Types)
	if err != nil {
		return
	}

	m.tasks[t.Client] = t
	return
}

func (m *Manager) Run() {
	for c, t := range m.tasks {
		del := run(t)
		if del {
			delete(m.tasks, c)
		}
	}
}

func run(t *Task) (del bool) {
	for _, typ := range t.Types {
		var (
			text string
			err  error
		)
		switch typ {
		case common.BiliSign:
			err = bili.SignIn(t.Client)

		case common.BiliBCount:
			err = bili.QueryBi(t.Client)

		case common.HPISign:
			err = hpi.SignIn(t.Client, t.ID)

		case common.STGSign:
			err = stg.SignIn(t.Client)
		}

		if err != nil {
			text = fmt.Sprintf("任务执行失败, id: %s, task: %s, err: %s",
				t.ID, common.TypeMap[typ], err)
			del = true
		} else {
			text = fmt.Sprintf("任务执行成功, id: %s, task: %s",
				t.ID, common.TypeMap[typ])
		}
		bot.Send(text)
		if del {
			return
		}
	}
	return
}

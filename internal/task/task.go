package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/robfig/cron/v3"

	"github.com/jdxj/sign/internal/pkg/bot"
	"github.com/jdxj/sign/internal/pkg/logger"
	"github.com/jdxj/sign/internal/storage"
	"github.com/jdxj/sign/internal/task/bili"
	"github.com/jdxj/sign/internal/task/common"
	"github.com/jdxj/sign/internal/task/hpi"
	"github.com/jdxj/sign/internal/task/stg"
	"github.com/jdxj/sign/internal/task/v2ex"
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
	_, _ = c.AddFunc("0 23 * * *", saveTasks)
}

type Task struct {
	ID     string       `json:"id"`
	Domain int          `json:"domain"`
	Types  []int        `json:"types"`
	Key    string       `json:"key"`
	Client *http.Client `json:"-"`
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

	case common.V2exDomain:
		t.Client, err = v2ex.Auth(t.Key)

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
			msg  string
			err  error
		)
		switch typ {
		case common.BiliSign:
			err = bili.SignIn(t.Client)

		case common.BiliBCount:
			msg, err = bili.QueryBi(t.Client)

		case common.HPISign:
			err = hpi.SignIn(t.Client, t.ID)

		case common.STGSign:
			err = stg.SignIn(t.Client)

		case common.V2exSign:
			err = v2ex.SignIn(t.Client)
		}

		if err != nil {
			text = fmt.Sprintf("任务执行失败, id: %s, task: %s, err: %s",
				t.ID, common.TypeMap[typ], err)
			del = true
		} else {
			text = fmt.Sprintf("任务执行成功, id: %s, task: %s",
				t.ID, common.TypeMap[typ])
			if len(msg) != 0 {
				text = fmt.Sprintf("%s, msg: %s", text, msg)
			}
		}

		bot.Send(text)

		if del {
			return
		}
	}
	return
}

func (m *Manager) Marshal() ([]byte, error) {
	tasks := make([]*Task, 0, len(m.tasks))
	for _, v := range m.tasks {
		req, _ := http.NewRequest("", "", nil)
		u := getHomeURL(v.Domain)
		cookies := v.Client.Jar.Cookies(u)
		for _, cookie := range cookies {
			req.AddCookie(cookie)
		}
		v.Key = req.Header.Get("Cookie")
		tasks = append(tasks, v)
	}
	return json.Marshal(tasks)
}

func getHomeURL(domain int) *url.URL {
	var (
		u    *url.URL
		home string
	)
	switch domain {
	case common.BiliDomain:
		home = bili.SignURL
	case common.HPIDomain:
		home = hpi.URL
	case common.STGDomain:
		home = stg.HomeURL
	case common.V2exDomain:
		home = v2ex.Home
	}
	u, _ = url.Parse(home)
	return u
}

func (m *Manager) Unmarshal(data []byte) error {
	var tasks []*Task
	err := json.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}

	for _, t := range tasks {
		err = m.Add(t)
		if err != nil {
			return fmt.Errorf("unmarshal task failed, id: %s, types: %v, err: %w",
				t.ID, t.Types, err)
		}
	}
	return nil
}

func saveTasks() {
	data, err := manager.Marshal()
	if err != nil {
		logger.Errorf("marshal tasks failed, err: %s", err)
		text := fmt.Sprintf("marshal tasks failed, err: %s", err)
		bot.Send(text)
		return
	}

	err = storage.Write(data)
	if err != nil {
		logger.Errorf("write data failed, err: %s", err)
		text := fmt.Sprintf("write data failed, err: %s", err)
		bot.Send(text)
		return
	}

	text := "save tasks success"
	bot.Send(text)
}

func RecoverTasks() {
	data, err := storage.Read()
	if err != nil {
		logger.Errorf("read data failed, err: %s", err)
		text := fmt.Sprintf("read data failed, err: %s", err)
		bot.Send(text)
		return
	}
	if len(data) == 0 {
		logger.Infof("have no tasks to recover")
		return
	}

	err = manager.Unmarshal(data)
	if err != nil {
		logger.Errorf("unmarshal task failed, err: %s", err)
		text := fmt.Sprintf("unmarshal task failed, err: %s", err)
		bot.Send(text)
		return
	}

	text := "recover tasks success"
	bot.Send(text)
}

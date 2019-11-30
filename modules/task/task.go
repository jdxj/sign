package task

import (
	"fmt"
	"math/rand"
	"sign/modules"
	"sign/utils/email"
	"sign/utils/log"
	"sync"
	"time"
)

var DefaultExe *Executor

func init() {
	DefaultExe = &Executor{}
}

type Executor struct {
	touchers []modules.Toucher

	locker sync.RWMutex
}

// AddTaskFromApi 从 http api 接收一个 toucher,
// 其使用一个 goroutine 启动任务.
// 注意: 目前没有对所启动 goroutine 进行管理, 所以是不安全的.
func (exe *Executor) AddTaskFromApi(tou modules.Toucher) error {
	if tou == nil {
		return fmt.Errorf("toucher is nil")
	}

	if !tou.Boot() {
		// 通过 api 接口创建的任务, 如果 boot 阶段就失败, 则直接向
		// api 接口返回错误, 不写入日志且不使用邮件通知
		return fmt.Errorf("boot fail")
	}

	go func() {
		tomSome := randTime()

		dur := tomSome.Sub(time.Now())
		timer := time.NewTimer(dur)
		defer timer.Stop()

		for {
			<-timer.C

			if !tou.Login() {
				log.MyLogger.Error("%s login fail: %s", log.Log_Task, tou.Name())
				email.SendEmail("签到失败通知", "task name: %s, stage: %s\n如果要重新签到, 请重新注册该任务", tou.Name(), "Login()")
				return
			}
			if !tou.Sign() {
				log.MyLogger.Error("%s sign fail: %s", log.Log_Task, tou.Name())
				email.SendEmail("签到失败通知", "task name: %s, stage: %s\n如果要重新签到, 请重新注册该任务", tou.Name(), "Sign()")
				return
			}

			tomSome = randTime()
			email.SendEmail("签到执行预通知", "签到任务: [%s] 将在 %s 时刻执行", tou.Name(), tomSome.Format(time.RFC1123))

			dur = tomSome.Sub(time.Now())
			timer.Reset(dur)
		}
	}()
	return nil
}

// randTime 返回明天的某个时刻
func randTime() time.Time {
	now := time.Now()
	r := rand.New(rand.NewSource(now.UnixNano()))

	tomNow := now.Add(24 * time.Hour)
	tom0AM := time.Date(tomNow.Year(), tomNow.Month(), tomNow.Day(), 0, 0, 0, 0, tomNow.Location())
	tom830AM := tom0AM.Add(8 * time.Hour).Add(30 * time.Minute)

	inc := r.Intn(12 * 60 * 60)
	tomSome := tom830AM.Add(time.Duration(inc) * time.Second)
	return tomSome
}

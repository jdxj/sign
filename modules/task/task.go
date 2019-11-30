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
		// api 接口返回错误, 不写入日志
		return fmt.Errorf("boot fail")
	}

	go func() {
		tomSome := randTime()

		dur := tomSome.Sub(time.Now())
		timer := time.NewTimer(dur)
		defer timer.Stop()

		msg := &email.Msg{
			To: tou.Email(),
		}

		msg.Subject = email.SignStart
		msg.Content = fmt.Sprintf("标记: %s, 下一次签到时间: %s", tou.Name(), tomSome.Format(email.TimeFormat))
		email.SendEmail(msg)

		for {
			<-timer.C

			if !tou.Login() {
				log.MyLogger.Error("%s login fail: %s", log.Log_Task, tou.Name())

				msg.Subject = email.SignFailed
				msg.Content = fmt.Sprintf("标记: %s, 阶段: %s", tou.Name(), "登录阶段")
				email.SendEmail(msg)
				return
			}
			if !tou.Sign() {
				log.MyLogger.Error("%s sign fail: %s", log.Log_Task, tou.Name())

				msg.Subject = email.SignFailed
				msg.Content = fmt.Sprintf("标记: %s, 阶段: %s", tou.Name(), "签到阶段")
				email.SendEmail(msg)
				return
			}

			tomSome = randTime()
			dur = tomSome.Sub(time.Now())
			timer.Reset(dur)

			msg.Subject = email.SignSuccess
			msg.Content = fmt.Sprintf("标记: %s, 下一次签到时间: %s", tou.Name(), tomSome.Format(email.TimeFormat))
			email.SendEmail(msg)
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

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

// todo: 将来会被弃用
func (exe *Executor) Run() {
	timer := time.NewTimer(time.Hour)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 无限循环
	for {
		// 现在时刻
		now := time.Now()
		// 明天这个时候
		tomNow := now.Add(24 * time.Hour)
		// 明天0点
		tom0AM := time.Date(tomNow.Year(), tomNow.Month(), tomNow.Day(), 0, 0, 0, 0, tomNow.Location())
		tom830AM := tom0AM.Add(8 * time.Hour).Add(30 * time.Minute)
		// 明天随便几点 8:30~20:30
		inc := time.Duration(r.Intn(12 * 60 * 60))
		tomSome := tom830AM.Add(inc * time.Second)
		// 下次签到时延
		dur := tomSome.Sub(now)

		timer.Reset(dur)
		log.MyLogger.Debug("等待时间到达: %s", tomSome)
		email.SendEmail("签到执行预通知", "签到任务将在 %s 时刻执行", tomSome.Format(time.RFC1123))

		select {
		case <-timer.C:
			exe.execute()
		}

		email.SendEmail("签到执行完成通知", "签到任务在 %s 时刻完成", tomSome.Format(time.RFC1123))
		log.MyLogger.Debug("%s", "本次每日任务完成...")
	}
}

// todo: 将来会被弃用
// execute 执行容器中的签到任务
func (exe *Executor) execute() {
	exe.locker.RLock()
	defer exe.locker.RUnlock()

	for _, toucher := range exe.touchers {
		if !toucher.Login() {
			log.MyLogger.Error("%s login fail: %s", log.Log_Task, toucher.Name())
			email.SendEmail("签到失败通知", "section: %s, stage: %s", toucher.Name(), "Login()")

			continue
		}
		if !toucher.Sign() {
			log.MyLogger.Error("%s sign fail: %s", log.Log_Task, toucher.Name())
			email.SendEmail("签到失败通知", "section: %s, stage: %s", toucher.Name(), "Sign()")
		}
	}
}

// todo: 将来会被弃用
// DebugRun 用于测试 task 的运行情况,
// 该方法立即执行容器中的 task.
func (exe *Executor) DebugRun() {
	exe.execute()
}

// todo: 将来会被弃用
// AddTaskAsync 向容器中添加任务, 非阻塞方式.
func (exe *Executor) AddTaskAsync(touchers ...modules.Toucher) {
	go func() {
		exe.locker.Lock()
		defer exe.locker.Unlock()

		for i, toucher := range touchers {
			if !toucher.Boot() {
				log.MyLogger.Warn("%s boot %s fail", log.Log_Task, toucher.Name())
				email.SendEmail("签到失败通知", "section: %s, stage: %s", toucher.Name(), "Boot()")
				continue
			}

			exe.touchers = append(exe.touchers, touchers[i])
		}
	}()
}

// todo: 将来会被弃用
// AddTaskSync 向容器中添加任务, 阻塞方式.
func (exe *Executor) AddTaskSync(touchers ...modules.Toucher) {
	exe.locker.Lock()
	defer exe.locker.Unlock()

	for i, toucher := range touchers {
		if !toucher.Boot() {
			log.MyLogger.Warn("%s boot %s fail", log.Log_Task, toucher.Name())
			email.SendEmail("签到失败通知", "section: %s, stage: %s", toucher.Name(), "Boot()")
			continue
		}

		exe.touchers = append(exe.touchers, touchers[i])
	}
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
		email.SendEmail("签到执行预通知", "签到任务: [%s] 将在 %s 时刻执行", tou.Name(), tomSome.Format(time.RFC1123))

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

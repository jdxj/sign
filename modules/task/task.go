package task

import (
	"math/rand"
	"sign/modules"
	"sign/utils/log"
	"sync"
	"time"
)

type Executor struct {
	touchers []modules.Toucher

	locker sync.RWMutex
}

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
		tomZero := time.Date(tomNow.Year(), tomNow.Month(), tomNow.Day(), 0, 0, 0, 0, tomNow.Location())
		// 明天随便几点
		inc := time.Duration(r.Intn(24 * 60 * 60))
		tomSome := tomZero.Add(inc * time.Second)
		// 下次签到时延
		dur := tomSome.Sub(now)

		timer.Reset(dur)
		log.MyLogger.Debug("等待时间到达: %s", tomSome)

		select {
		case <-timer.C:
			exe.execute()
		}

		log.MyLogger.Debug("%s", "本次每日任务完成...")
	}
}

// execute 执行容器中的签到任务
func (exe *Executor) execute() {
	exe.locker.RLock()
	defer exe.locker.RUnlock()

	for _, toucher := range exe.touchers {
		if !toucher.Login() {
			log.MyLogger.Error("%s login fail: %s", log.Log_Task, toucher.Name())
		}
		if !toucher.Sign() {
			log.MyLogger.Error("%s sign fail: %s", log.Log_Task, toucher.Name())
		}
	}
}

// DebugRun 用于测试 task 的运行情况,
// 该方法立即执行容器中的 task.
func (exe *Executor) DebugRun() {
	exe.execute()
}

// AddTaskAsync 向容器中添加任务, 非阻塞方式.
func (exe *Executor) AddTaskAsync(touchers ...modules.Toucher) {
	go func() {
		exe.locker.Lock()
		defer exe.locker.Unlock()

		for i, toucher := range touchers {
			if !toucher.Boot() {
				log.MyLogger.Warn("%s boot %s fail", log.Log_Task, toucher.Name())
				continue
			}

			exe.touchers = append(exe.touchers, touchers[i])
		}
	}()
}

// AddTaskSync 向容器中添加任务, 阻塞方式.
func (exe *Executor) AddTaskSync(touchers ...modules.Toucher) {
	exe.locker.Lock()
	defer exe.locker.Unlock()

	for i, toucher := range touchers {
		if !toucher.Boot() {
			log.MyLogger.Warn("%s boot %s fail", log.Log_Task, toucher.Name())
			continue
		}

		exe.touchers = append(exe.touchers, touchers[i])
	}
}

package utils

import (
	"math/rand"
	"time"
)

// DailyRandTimeExec 用于每日签到, 做了随机
func DailyRandTimeExec(prefix string, f func()) {
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
		LogPrintln(prefix, "等待时间到达:", tomSome)
		select {
		case <-timer.C:
			f()
		}
		LogPrintln(prefix, "本次每日任务完成...")
	}
}

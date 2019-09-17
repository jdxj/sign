package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// DailyRandTimeExec 用于每日签到, 做了随机
func DailyRandTimeExec(f func()) {
	timer := time.NewTimer(0)
	// 无限循环
	for {
		// 现在时刻
		now := time.Now()
		// 明天这个时候
		tomNow := now.Add(24 * time.Hour)
		// 明天0点
		tomZero := time.Date(tomNow.Year(), tomNow.Month(), tomNow.Day(), 0, 0, 0, 0, tomNow.Location())
		// 明天随便几点
		inc := time.Duration(rand.Intn(24 * 60 * 60))
		tomSome := tomZero.Add(inc * time.Second)
		// 下次签到时延
		dur := tomSome.Sub(now)

		timer.Reset(dur)
		fmt.Println("等待时间到达...")
		select {
		case <-timer.C:
			f()
		}
		fmt.Println("本次每日任务完成...")
	}
}

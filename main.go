package main

import (
	sgc "sign/studygolang.com"
	"sign/utils"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		utils.DailyRandTimeExec(utils.Log_StudyGolang, sgc.Start)
		// 目前不会退出, 只是用于阻塞
		wg.Done()
	}()

	wg.Wait()
}

package main

import (
	pic "sign/58pic.com"
	bili "sign/bilibili.com"
	sgc "sign/studygolang.com"
	"sign/utils"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)

	go func() {
		utils.DailyRandTimeExec(utils.Log_StudyGolang, sgc.Start)
		// 目前不会退出, 只是用于阻塞
		wg.Done()
	}()

	go func() {
		utils.DailyRandTimeExec(utils.Log_Bilibili, bili.Start)
		wg.Done()
	}()

	go func() {
		utils.DailyRandTimeExec(utils.Log_58pic, pic.Start)
		wg.Done()
	}()

	wg.Wait()
}

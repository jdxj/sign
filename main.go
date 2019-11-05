package main

import (
	"sign/cmd"
	"sign/modules"
	"sign/modules/task"
	"sign/utils"
	"sign/utils/conf"
)

//import (
//	pic "sign/58pic.com"
//	bili "sign/bilibili.com"
//	sgc "sign/studygolang.com"
//	"sign/utils"
//	"sync"
//)
//
//func main() {
//	wg := &sync.WaitGroup{}
//	wg.Add(3)
//
//	go func() {
//		utils.DailyRandTimeExec(utils.Log_StudyGolang, sgc.Start)
//		// 目前不会退出, 只是用于阻塞
//		wg.Done()
//	}()
//
//	go func() {
//		utils.DailyRandTimeExec(utils.Log_Bilibili, bili.Start)
//		wg.Done()
//	}()
//
//	go func() {
//		utils.DailyRandTimeExec(utils.Log_58pic, pic.Start)
//		wg.Done()
//	}()
//
//	wg.Wait()
//}

func main() {
	utils.MyLogger.Debug("%s sections' len: %d", utils.Log_Main, len(conf.Conf.Sections()))

	var touchers []modules.Toucher
	for _, sec := range conf.Conf.Sections() {
		if sec.Name() == "email" || sec.Name() == "DEFAULT" {
			utils.MyLogger.Warn("%s jump over %s section", utils.Log_Main, sec.Name())
			continue
		}

		toucher, err := cmd.NewToucher(sec)
		if err != nil {
			utils.MyLogger.Error("%s %s, section name: %s, %s", utils.Log_Main, "create toucher fail", sec.Name(), err)
			continue
		}

		touchers = append(touchers, toucher)
	}

	if len(touchers) == 0 {
		utils.MyLogger.Warn("%s %s", utils.Log_Main, "add 0 touchers")
	}

	exe := &task.Executor{}
	exe.AddTaskSync(touchers...)
	exe.Run()
	//exe.DebugRun()
}

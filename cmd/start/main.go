package main

import (
	"github.com/jdxj/sign/configs"
	"github.com/jdxj/sign/internal/bot"
	"github.com/jdxj/sign/pkg/task"
)

func main() {
	root := configs.ReadConfigs("configs.yaml")
	bot.Init(root.Bot.Token, root.Bot.ChatID)
	task.Run()
}

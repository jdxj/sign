package service

import (
	pb "github.com/jdxj/sign/internal/proto/task"
	"github.com/jdxj/sign/internal/task/service/executor"
	"github.com/jdxj/sign/internal/task/service/executor/bilibili"
	"github.com/jdxj/sign/internal/task/service/executor/bwh"
	custom_message "github.com/jdxj/sign/internal/task/service/executor/custom-message"
	"github.com/jdxj/sign/internal/task/service/executor/evo"
	"github.com/jdxj/sign/internal/task/service/executor/github"
	"github.com/jdxj/sign/internal/task/service/executor/juejin"
	"github.com/jdxj/sign/internal/task/service/executor/stg"
	"github.com/jdxj/sign/internal/task/service/executor/v2ex"
)

var (
	executors = map[string]Executor{
		pb.Kind_MOCK.String(): &executor.MockExecutor{},

		pb.Kind_BILIBILI_SIGN_IN.String(): &bilibili.SignIn{},
		pb.Kind_BILIBILI_B_COUNT.String(): &bilibili.Bi{},

		pb.Kind_EVOLUTION_RELEASE.String(): &evo.Updater{},

		pb.Kind_GITHUB_RELEASE.String(): &github.Release{},

		pb.Kind_JUEJIN_SIGN_IN.String():  &juejin.SignIn{},
		pb.Kind_JUEJIN_POINT.String():    &juejin.Point{},
		pb.Kind_JUEJIN_COUNT.String():    &juejin.Count{},
		pb.Kind_JUEJIN_CALENDAR.String(): &juejin.Calendar{},

		pb.Kind_STG_SIGN_IN.String(): &stg.SignIn{},

		pb.Kind_V2EX_SIGN_IN.String(): &v2ex.SignIn{},

		pb.Kind_CUSTOM_MESSAGE.String(): &custom_message.CustomMessage{},

		pb.Kind_BWH_NETWORK.String(): &bwh.NetworkMeter{},
		pb.Kind_BWH_CPU.String():     &bwh.CPUMeter{},
	}
)

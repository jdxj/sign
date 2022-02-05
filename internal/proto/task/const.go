package task

import (
	"google.golang.org/protobuf/proto"
)

const (
	ServiceName = "task"
	Topic       = "task-dispatch"
	Queue       = "consume-task"
)

func GetParamByKind(kind string) (msg proto.Message) {
	switch kind {
	case Kind_BILIBILI_SIGN_IN.String(), Kind_BILIBILI_B_COUNT.String():
		msg = &BiLiBiLi{}
	case Kind_STG_SIGN_IN.String():
		msg = &STG{}
	case Kind_V2EX_SIGN_IN.String():
		msg = &V2Ex{}
	case Kind_EVOLUTION_RELEASE.String():
		msg = &Evolution{}
	case Kind_GITHUB_RELEASE.String():
		msg = &GithubRelease{}
	case Kind_JUEJIN_SIGN_IN.String(), Kind_JUEJIN_COUNT.String(),
		Kind_JUEJIN_POINT.String(), Kind_JUEJIN_CALENDAR.String():
		msg = &JueJin{}
	case Kind_CUSTOM_MESSAGE.String():
		msg = &CustomMessage{}
	case Kind_BWH_NETWORK.String(), Kind_BWH_CPU.String():
		msg = &BWH{}
	}
	return
}

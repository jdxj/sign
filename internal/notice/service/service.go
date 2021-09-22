package service

import (
	"context"
	"errors"
	"fmt"

	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/proto/notice"
	"github.com/jdxj/sign/internal/proto/user"
)

var (
	ErrInvalidConfig = errors.New("invalid config")
)

func New(conf config.Bot) *Service {
	if conf.Token == "" || conf.ChatID == 0 {
		panic(ErrInvalidConfig)
	}
	srv := &Service{
		chatID: conf.ChatID,
	}
	client, err := bot.NewBotAPI(conf.Token)
	if err != nil {
		panic(err)
	}
	srv.botClient = client

	rpc.NewClient(user.ServiceName, func(cc *grpc.ClientConn) {
		srv.userClient = user.NewUserServiceClient(cc)
	})
	return srv
}

type Service struct {
	notice.UnimplementedNoticeServiceServer

	botClient  *bot.BotAPI
	chatID     int64
	userClient user.UserServiceClient
}

func (srv *Service) SendMessage(ctx context.Context, req *notice.SendMessageReq) (*emptypb.Empty, error) {
	if req.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}
	if req.Text == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty text")
	}

	userInfo, err := srv.userClient.GetUser(ctx, &user.GetUserReq{UserID: req.UserID})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "get user failed: %s", err)
	}
	text := fmt.Sprintf("%s: %s", userInfo.Nickname, req.Text)

	msgConf := bot.NewMessage(srv.chatID, text)
	_, err = srv.botClient.Send(msgConf)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "send message failed: %s", err)
	}
	return &emptypb.Empty{}, nil
}

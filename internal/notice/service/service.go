package service

import (
	"context"
	"errors"
	"log"

	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/pkg/config"
	pb "github.com/jdxj/sign/internal/proto/notice"
	"github.com/jdxj/sign/internal/proto/user"
)

var (
	ErrInvalidConfig = errors.New("invalid config")
)

func New(conf config.Bot) *Service {
	if conf.Token == "" {
		panic(ErrInvalidConfig)
	}
	srv := &Service{}
	log.Printf("new bot api...")
	client, err := bot.NewBotAPI(conf.Token)
	if err != nil {
		panic(err)
	}
	srv.botClient = client
	return srv
}

type Service struct {
	botClient *bot.BotAPI
}

func (srv *Service) SendNotice(ctx context.Context, req *pb.SendNoticeRequest, rsp *emptypb.Empty) error {
	if req.GetUserId() == 0 {
		return status.Errorf(codes.InvalidArgument, "empty user id")
	}

	userRsp, err := UserService.GetUser(ctx, &user.GetUserRequest{UserID: req.GetUserId()})
	if err != nil {
		return status.Errorf(codes.Internal, "GetUser: %s", err)
	}

	// todo: 配置化
	contact := userRsp.GetUser().GetContact().GetTelegram()
	if contact == 0 {
		return nil
	}

	msgConf := bot.NewMessage(contact, req.GetContent())
	_, err = srv.botClient.Send(msgConf)
	if err != nil {
		return status.Errorf(codes.Internal, "Send: %s", err)
	}
	return nil
}

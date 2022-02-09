package service

import (
	"context"
	"errors"

	bot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/notice/cache"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/logger"
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

	logger.Debugf("new bot api...")
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

func (srv *Service) SendNotice(ctx context.Context, req *pb.SendNoticeRequest, _ *emptypb.Empty) error {
	if req.GetUserId() == 0 {
		return status.Errorf(codes.InvalidArgument, "empty user id")
	}

	// todo: 配置化
	contact := cache.GetUserTelegram(ctx, req.GetUserId())
	if contact == 0 {
		userRsp, err := UserService.GetUser(ctx, &user.GetUserRequest{UserID: req.GetUserId()})
		if err != nil {
			return status.Errorf(codes.Internal, "GetUser: %s", err)
		}

		contact = userRsp.GetUser().GetContact().GetTelegram()
		if contact == 0 {
			return nil
		}
		cache.SetUserTelegram(ctx, req.GetUserId(), contact)
	}

	msgConf := bot.NewMessage(contact, req.GetContent())
	_, err := srv.botClient.Send(msgConf)
	if err != nil {
		logger.Errorf("Send: %s, userID: %d", err, req.GetUserId())
		return status.Errorf(codes.Internal, "Send: %s", err)
	}
	return nil
}

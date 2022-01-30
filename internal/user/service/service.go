package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/jdxj/sign/internal/pkg/db"
	"github.com/jdxj/sign/internal/pkg/util"
	pb "github.com/jdxj/sign/internal/proto/user"
	"github.com/jdxj/sign/internal/user/dao"
)

type Service struct {
}

func (srv *Service) AuthUser(ctx context.Context, req *pb.AuthUserRequest, rsp *pb.AuthUserResponse) error {
	if req.GetNickname() == "" || req.GetPassword() == "" {
		return status.Errorf(codes.InvalidArgument, "invalid params")
	}

	var row dao.User
	err := db.WithCtx(ctx).
		Where("nickname = ?", req.GetNickname()).
		Take(&row).
		Error
	if err != nil {
		return status.Errorf(codes.Internal, err.Error())
	}
	pass := util.WithSalt(req.GetPassword(), row.Salt)
	if pass == row.Password {
		rsp.Valid = true
		rsp.UserID = row.UserID
	}
	return nil
}

func (srv *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest, rsp *pb.CreateUserResponse) error {
	if req.GetUser().GetNickname() == "" || req.GetUser().GetPassword() == "" {
		return status.Errorf(codes.InvalidArgument, "invalid params")
	}

	salt := util.Salt()
	pass := req.GetUser().GetPassword()
	user := dao.User{
		Nickname: req.GetUser().GetNickname(),
		Password: util.WithSalt(pass, salt),
		Salt:     salt,
		Mail:     req.GetUser().GetContact().GetMail(),
		Telegram: req.GetUser().GetContact().GetTelegram(),
	}

	err := db.WithCtx(ctx).
		Create(&user).
		Error
	if err != nil {
		return status.Errorf(codes.Internal, "Create: %s", err)
	}
	rsp.UserID = user.UserID
	return nil
}

func (srv *Service) GetUser(ctx context.Context, req *pb.GetUserRequest, rsp *pb.GetUserResponse) error {
	if req.GetUserID() == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid params")
	}

	var row dao.User
	err := db.WithCtx(ctx).
		Where("user_id = ?", req.GetUserID()).
		Take(&row).
		Error
	if err != nil {
		return status.Errorf(codes.Internal, "Take: %s", err)
	}

	rsp.User = &pb.User{
		UserId:   row.UserID,
		Nickname: row.Nickname,
		Contact: &pb.Contact{
			Mail:     row.Mail,
			Telegram: row.Telegram,
		},
	}
	return nil
}

func (srv *Service) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest, _ *emptypb.Empty) error {
	user := req.GetUser()
	if user.GetUserId() == 0 {
		return status.Errorf(codes.InvalidArgument, "invalid params")
	}

	data := make(map[string]interface{})
	if user.GetNickname() != "" {
		data["nickname"] = user.GetNickname()
	}
	if user.GetPassword() != "" {
		salt := util.Salt()
		data["salt"] = salt
		data["password"] = util.WithSalt(user.GetPassword(), salt)
	}

	contact := user.GetContact()
	if contact.GetMail() != "" {
		data["mail"] = contact.GetMail()
	}
	if contact.GetTelegram() != 0 {
		data["telegram"] = contact.GetTelegram()
	}

	err := db.WithCtx(ctx).
		Table(dao.TableName).
		Where("user_id = ?", user.GetUserId()).
		Updates(data).
		Error
	if err != nil {
		return status.Errorf(codes.Internal, "Updates: %s", err)
	}
	return nil
}

package service

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/jdxj/sign/internal/pkg/util"
	"github.com/jdxj/sign/internal/proto/user"
	userDao "github.com/jdxj/sign/internal/user/dao/user"
)

type Service struct {
	user.UnimplementedUserServiceServer
}

func (srv *Service) CreateUser(ctx context.Context, req *user.CreateUserReq) (*user.CreateUserRsp, error) {
	if req.Nickname == "" || req.Password == "" {
		return nil, status.Errorf(codes.InvalidArgument, "empty params")
	}

	u := &userDao.User{
		Nickname: req.Nickname,
		Salt:     util.Salt(),
	}
	u.Password = util.WithSalt(req.Password, u.Salt)
	err := userDao.Insert(u)
	rsp := &user.CreateUserRsp{UserID: u.UserID}
	return rsp, err
}

func (srv *Service) GetUser(ctx context.Context, req *user.GetUserReq) (*user.GetUserRsp, error) {
	if req.UserID == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "empty user id")
	}

	u, err := userDao.FindOne(map[string]interface{}{"user_id = ?": req.UserID})
	if err != nil {
		return nil, err
	}
	rsp := &user.GetUserRsp{
		Nickname: u.Nickname,
		// todo: 是否返回
		Password: "",
	}
	return rsp, nil
}

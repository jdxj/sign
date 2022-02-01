package handler

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/jdxj/sign/internal/app/api"
	"github.com/jdxj/sign/internal/app/ref"
	"github.com/jdxj/sign/internal/pkg/logger"
	ser "github.com/jdxj/sign/internal/pkg/sign-error"
	"github.com/jdxj/sign/internal/proto/user"
)

type LoginReq struct {
	Nickname string `json:"nickname" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRsp struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func login(ctx context.Context, req *LoginReq) (*LoginRsp, error) {
	auRsp, err := ref.UserService.AuthUser(ctx, &user.AuthUserRequest{
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		logger.Errorf("AuthUser: %s", err)
		return nil, ser.New(ser.ErrRPCRequest, "认证失败")
	}
	if !auRsp.GetValid() {
		return nil, ser.New(ser.ErrAuthFailed, "认证失败")
	}

	rsp := &LoginRsp{
		UserID: auRsp.GetUserID(),
	}
	token, err := api.NewClaim(auRsp.GetUserID(), req.Nickname).Token()
	if err != nil {
		logger.Errorf("GenerateToken: %s", err)
		return nil, ser.New(ser.ErrInternal, "认证失败")
	}
	rsp.Token = token
	return rsp, nil
}

func Login(ctx *gin.Context) {
	req := &LoginReq{}
	api.Process(ctx, req, func(request *api.Request) (interface{}, error) {
		return login(ctx, req)
	})
}

type SignUpReq struct {
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Mail     string `json:"mail"`
	Telegram int64  `json:"telegram"`
}

type SignUpRsp struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func signUp(ctx context.Context, req *SignUpReq) (*SignUpRsp, error) {
	cuRsp, err := ref.UserService.CreateUser(ctx, &user.CreateUserRequest{User: &user.User{
		Nickname: req.Nickname,
		Password: req.Password,
		Contact: &user.Contact{
			Mail:     req.Mail,
			Telegram: req.Telegram,
		},
	}})
	if err != nil {
		return nil, ser.New(ser.ErrRPCRequest, "CreateUser: %s", err)
	}

	rsp := &SignUpRsp{
		UserID: cuRsp.GetUserID(),
	}
	token, err := api.NewClaim(cuRsp.UserID, req.Nickname).Token()
	if err != nil {
		return nil, ser.New(ser.ErrInternal, "获取 token 失败: %s", err)
	}
	rsp.Token = token
	return rsp, nil
}

func SignUp(ctx *gin.Context) {
	req := &SignUpReq{}
	api.Process(ctx, req, func(request *api.Request) (interface{}, error) {
		return signUp(ctx, req)
	})
}

type UpdateUserReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
	Telegram int64  `json:"telegram"`
}

func updateUser(ctx context.Context, req *UpdateUserReq, userID int64) error {
	_, err := ref.UserService.UpdateUser(ctx, &user.UpdateUserRequest{User: &user.User{
		UserId:   userID,
		Nickname: req.Nickname,
		Password: req.Password,
		Contact: &user.Contact{
			Mail:     req.Mail,
			Telegram: req.Telegram,
		},
	}})
	if err != nil {
		return ser.New(ser.ErrRPCRequest, "UpdateUser: %s", err)
	}
	return nil
}

func UpdateUser(ctx *gin.Context) {
	req := &UpdateUserReq{}
	api.ProcessCheckToken(ctx, req, func(request *api.Request) (interface{}, error) {
		return nil, updateUser(ctx, req, request.Claim.UserID)
	})
}

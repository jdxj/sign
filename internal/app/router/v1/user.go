package v1

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
	Nickname string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
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
		return nil, ser.Wrap(ser.ErrRPCCall, err, "AuthUser")
	}
	if !auRsp.GetValid() {
		return nil, ser.New(ser.ErrLogin, "invalid nickname or password")
	}

	rsp := &LoginRsp{
		UserID: auRsp.GetUserID(),
	}
	rsp.Token, err = api.NewSignClaim(auRsp.GetUserID(), req.Nickname).Token()
	if err != nil {
		logger.Errorf("GenerateToken: %s", err)
		return nil, ser.Wrap(ser.ErrUnknown, err, "NewSignClaim")
	}
	return rsp, nil
}

func Login(ctx *gin.Context) {
	req := &LoginReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		api.Respond(ctx, nil, ser.Wrap(ser.ErrBindRequest, err, "Login"))
		return
	}
	data, err := login(ctx, req)
	api.Respond(ctx, data, err)
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
		return nil, ser.Wrap(ser.ErrRPCCall, err, "CreateUser")
	}

	rsp := &SignUpRsp{
		UserID: cuRsp.GetUserID(),
	}
	rsp.Token, err = api.NewSignClaim(cuRsp.UserID, req.Nickname).Token()
	if err != nil {
		return nil, ser.Wrap(ser.ErrUnknown, err, "NewSignClaim")
	}
	return rsp, nil
}

func SignUp(ctx *gin.Context) {
	req := &SignUpReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		api.Respond(ctx, nil, ser.Wrap(ser.ErrBindRequest, err, "SignUp"))
		return
	}
	data, err := signUp(ctx, req)
	api.Respond(ctx, data, err)
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
		return ser.Wrap(ser.ErrRPCCall, err, "UpdateUser")
	}
	return nil
}

func UpdateUser(ctx *gin.Context) {
	req := &UpdateUserReq{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		api.Respond(ctx, nil, ser.Wrap(ser.ErrBindRequest, err, "UpdateUser"))
		return
	}
	err = updateUser(ctx, req, getUserID(ctx))
	api.Respond(ctx, nil, err)
}

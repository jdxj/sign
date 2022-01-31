package apiserver

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/rpc"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/code"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/proto/user"
)

var (
	ErrParseToken = errors.New("parse token")
)

var (
	JwtKey string

	UserClient   user.UserServiceClient
	CronClient   crontab.CrontabServiceClient
	SecretClient secret.SecretServiceClient
)

func Init(conf config.APIServer) {
	JwtKey = conf.Key

	rpc.NewClient(user.ServiceName, func(cc *grpc.ClientConn) {
		UserClient = user.NewUserServiceClient(cc)
	})
	rpc.NewClient(crontab.ServiceName, func(cc *grpc.ClientConn) {
		CronClient = crontab.NewCrontabServiceClient(cc)
	})
	rpc.NewClient(secret.ServiceName, func(cc *grpc.ClientConn) {
		SecretClient = secret.NewSecretServiceClient(cc)
	})
}

func NewClaim() *Claim {
	stdClaim := jwt.StandardClaims{
		Audience:  "",
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		Id:        "",
		IssuedAt:  0,
		Issuer:    "apiserver",
		NotBefore: 0,
		Subject:   "",
	}
	claim := &Claim{
		StandardClaims: stdClaim,
	}
	return claim
}

type Claim struct {
	jwt.StandardClaims

	UserID   int64
	Nickname string
}

func GenerateToken(claim *Claim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(JwtKey))
}

func CheckToken(sign string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(sign, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*Claim)
	if !ok || !token.Valid {
		return nil, ErrParseToken
	}
	return claim, nil
}

func TimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

type Request struct {
	Token string          `json:"token"`
	Data  json.RawMessage `json:"data"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Reply(ctx *gin.Context, code int, msg string, data interface{}) {
	rsp := &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, rsp)
}

type LoginReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type LoginRsp struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func Login(ctx *gin.Context) {
	req := &LoginReq{}
	err := ctx.Bind(req)
	if err != nil {
		Reply(ctx, code.ErrBindReqFailed, err.Error(), nil)
		return
	}

	authRsp, err := UserClient.AuthUser(ctx, &user.AuthUserReq{
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		Reply(ctx, code.ErrRPCRequest, err.Error(), nil)
		return
	}
	if !authRsp.Valid {
		Reply(ctx, code.ErrAuthFailed, "invalid nickname or password", nil)
		return
	}

	claim := NewClaim()
	claim.UserID = authRsp.UserID
	claim.Nickname = req.Nickname

	rsp := &LoginRsp{
		UserID: authRsp.UserID,
	}
	rsp.Token, err = GenerateToken(claim)
	if err != nil {
		Reply(ctx, code.ErrInternal, err.Error(), nil)
		return
	}
	Reply(ctx, 0, "", rsp)
}

const (
	KeyClaim = "claim"
	KeyData  = "data"
)

func Auth(ctx *gin.Context) {
	req := &Request{}
	err := ctx.Bind(req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, &Response{
			Code:    code.ErrBindReqFailed,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	claim, err := CheckToken(req.Token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, &Response{
			Code:    code.ErrAuthFailed,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	ctx.Set(KeyClaim, claim)
	ctx.Set(KeyData, req.Data)
}

func Handle(ctx *gin.Context, req interface{}, f func(context.Context) (interface{}, error)) {
	data, _ := ctx.Get(KeyData)
	rawMsg, _ := data.(json.RawMessage)
	err := json.Unmarshal(rawMsg, req)
	if err != nil {
		Reply(ctx, code.ErrInvalidParam, err.Error(), nil)
		return
	}

	tCtx, cancel := TimeoutContext()
	defer cancel()

	data, err = f(tCtx)
	if err != nil {
		Reply(ctx, code.ErrHandle, err.Error(), nil)
		return
	}
	Reply(ctx, 0, "", data)
}

type SignUpReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type SignUpRsp struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

func SignUp(ctx *gin.Context) {
	req := &SignUpReq{}
	err := ctx.Bind(req)
	if err != nil {
		Reply(ctx, code.ErrBindReqFailed, err.Error(), nil)
		return
	}

	tCtx, cancel := TimeoutContext()
	defer cancel()

	createRsp, err := UserClient.CreateUser(tCtx, &user.CreateUserReq{
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		Reply(ctx, code.ErrInternal, err.Error(), nil)
		return
	}

	claim := NewClaim()
	claim.UserID = createRsp.UserID
	claim.Nickname = req.Nickname
	rsp := &SignUpRsp{
		UserID: createRsp.UserID,
	}
	rsp.Token, err = GenerateToken(claim)
	if err != nil {
		Reply(ctx, code.ErrInternal, err.Error(), nil)
		return
	}
	Reply(ctx, 0, "", rsp)
}

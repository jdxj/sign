package comm

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"

	"github.com/jdxj/sign/internal/pkg/code"
	"github.com/jdxj/sign/internal/pkg/config"
	"github.com/jdxj/sign/internal/pkg/rpc"
	"github.com/jdxj/sign/internal/proto/user"
)

var (
	ErrInvalidKey = errors.New("invalid key")
)

var (
	cfg config.APIServer

	UserClient user.UserServiceClient
)

func Init(conf config.APIServer) {
	cfg = conf
	if cfg.Key == "" {
		panic(ErrInvalidKey)
	}

	rpc.NewClient(user.ServiceName, func(cc *grpc.ClientConn) {
		UserClient = user.NewUserServiceClient(cc)
	})
}

type Request struct {
	Token string          `json:"token"`
	Data  json.RawMessage `json:"data"`
}

func NewResponse(code int, msg string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Message: msg,
		Data:    data,
	}
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type AddTaskReq struct {
	ID     string `json:"id"`
	Domain int    `json:"domain"`
	Type   []int  `json:"type"`
	Key    string `json:"key"`
}

func UnmarshalRequest(ctx *gin.Context, req interface{}) error {
	data := ctx.GetString("data")
	err := json.Unmarshal([]byte(data), req)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			NewResponse(code.ErrBindReqFailed, "unmarshal request failed", nil),
		)
	}
	return err
}

type Claim struct {
	UserID   int64
	Nickname string
	jwt.StandardClaims
}

func GenerateToken(claim *Claim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	key := []byte(cfg.Key)
	return token.SignedString(key)
}

func CheckToken(sign string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(sign, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		key := []byte(cfg.Key)
		return key, nil
	})
	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*Claim)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("parse token failed")
	}
	return claim, nil
}

func TimeoutContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), 5*time.Second)
}

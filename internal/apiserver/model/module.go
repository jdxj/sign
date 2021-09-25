package model

import (
	"context"
	"fmt"

	"github.com/dgrijalva/jwt-go"

	"github.com/jdxj/sign/internal/apiserver/comm"
	"github.com/jdxj/sign/internal/proto/user"
)

type LoginReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

func GenerateToken(ctx context.Context, req *LoginReq) (*LoginResp, error) {
	rsp, err := comm.UserClient.AuthUser(ctx, &user.AuthUserReq{
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	if !rsp.Valid {
		return nil, fmt.Errorf("invalid nickname or password")
	}

	claim := &comm.Claim{
		UserID:   rsp.UserID,
		Nickname: req.Nickname,
		StandardClaims: jwt.StandardClaims{
			Issuer: "apiserver",
		},
	}
	resp := &LoginResp{}
	resp.Token, err = comm.GenerateToken(claim)
	return resp, err
}

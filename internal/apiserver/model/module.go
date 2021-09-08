package model

import (
	"github.com/dgrijalva/jwt-go"

	"github.com/jdxj/sign/internal/apiserver/comm"
	"github.com/jdxj/sign/internal/apiserver/dao/user"
)

type TestModule struct {
	Nickname string `json:"nickname"`
}

type LoginReq struct {
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

func GenerateToken(req *LoginReq) (*LoginResp, error) {
	u, err := user.Find(req.Nickname, req.Password)
	if err != nil {
		return nil, err
	}

	claim := &comm.Claim{
		UserID:   u.UserID,
		Nickname: u.Nickname,
		StandardClaims: jwt.StandardClaims{
			Issuer: "apiserver",
		},
	}
	resp := &LoginResp{}
	resp.Token, err = comm.GenerateToken(claim)
	return resp, err
}

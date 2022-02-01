package api

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	JwtKey string
)

func Init(jwtKey string) {
	JwtKey = jwtKey
}

var (
	ErrParseToken = errors.New("parse token")
)

func NewClaim(userID int64, nickname string) *Claim {
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
		UserID:         userID,
		Nickname:       nickname,
	}
	return claim
}

type Claim struct {
	jwt.StandardClaims

	UserID   int64
	Nickname string
}

func (c *Claim) Token() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(JwtKey))
}

func NewClaimFromToken(sign string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(sign, &Claim{}, func(*jwt.Token) (interface{}, error) {
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

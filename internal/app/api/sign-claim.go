package api

import (
	"errors"
	"math/rand"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
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

func NewSignClaim(userID int64, nickname string) *SignClaim {
	now := time.Now()
	sc := jwt.RegisteredClaims{
		Issuer:    "sign",
		Subject:   "app",
		ExpiresAt: jwt.NewNumericDate(now.AddDate(0, 0, 1)),
		NotBefore: jwt.NewNumericDate(now),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        strconv.Itoa(rand.Int()),
	}
	claim := &SignClaim{
		RegisteredClaims: sc,
		UserID:           userID,
		Nickname:         nickname,
	}
	return claim
}

type SignClaim struct {
	jwt.RegisteredClaims

	UserID   int64
	Nickname string
}

func (c *SignClaim) Token() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(JwtKey))
}

func NewSignClaimFromToken(sign string) (*SignClaim, error) {
	token, err := jwt.ParseWithClaims(sign, &SignClaim{}, func(*jwt.Token) (interface{}, error) {
		return []byte(JwtKey), nil
	})
	if err != nil {
		return nil, err
	}

	claim, ok := token.Claims.(*SignClaim)
	if !ok || !token.Valid {
		return nil, ErrParseToken
	}
	return claim, nil
}

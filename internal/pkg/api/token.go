package api

import (
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Claim struct {
	jwt.StandardClaims

	UserID   int64
	Nickname string
}

func GenerateToken(key string, claim *Claim) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(key))
}

func CheckToken(key, sign string) (*Claim, error) {
	token, err := jwt.ParseWithClaims(sign, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
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

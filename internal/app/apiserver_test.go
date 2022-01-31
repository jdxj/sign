package app

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/jdxj/sign/internal/app/api"
	"github.com/jdxj/sign/internal/app/handler"
)

func TestGenerateToken(t *testing.T) {
	api.JwtKey = "jdxj"
	claim := &handler.Claim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Second).Unix(),
		},
		UserID:   1,
		Nickname: "test",
	}
	token, err := handler.GenerateToken(claim)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	time.Sleep(5 * time.Second)

	claim, err = handler.CheckToken(token)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", claim)
}

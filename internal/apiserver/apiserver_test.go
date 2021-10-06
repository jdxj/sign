package apiserver

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestGenerateToken(t *testing.T) {
	JwtKey = "jdxj"
	claim := &Claim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(3 * time.Second).Unix(),
		},
		UserID:   1,
		Nickname: "test",
	}
	token, err := GenerateToken(claim)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	time.Sleep(5 * time.Second)

	claim, err = CheckToken(token)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", claim)
}

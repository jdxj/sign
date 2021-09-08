package comm

import (
	"fmt"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestCheckToken(t *testing.T) {
	claim := &Claim{
		UserID:   1,
		Nickname: "jdxj",
		StandardClaims: jwt.StandardClaims{
			Issuer: "apiserver",
		},
	}

	sign, err := GenerateToken(claim)
	if err != nil {
		t.Fatalf("%s\n", err)
	}

	claim, err = CheckToken(sign)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", *claim)
}

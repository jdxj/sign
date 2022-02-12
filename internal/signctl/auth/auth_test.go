package auth

import (
	"fmt"
	"net/url"
	"testing"
)

func TestJoin(t *testing.T) {
	u, err := url.Parse("https://example.com")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("u1: %s\n", u.String())
	fmt.Printf("u2: %s\n", u.Path)
	fmt.Printf("u3: %s\n", u.RawPath)

	//u.Path = path.Join(u.Path, "abc")
	u.Path = "/abc"
	fmt.Printf("u4: %s\n", u.String())
}

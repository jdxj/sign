package task

import (
	"fmt"
	"testing"
)

type A struct {
	Name string
	Addr string
}

func TestPopulateStruct(t *testing.T) {
	key := "name=jdxj;addr=earth"
	a := &A{}
	err := PopulateStruct(key, a)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%+v\n", a)
}

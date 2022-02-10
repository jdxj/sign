package main

import (
	"fmt"
	"go/build"
	"testing"
)

func TestBuild(t *testing.T) {
	pkg, err := build.ImportDir(".", 0)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("pkg: %+v\n", pkg)
}

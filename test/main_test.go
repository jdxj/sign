package main

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {
	str := "bearer abc"
	var res string
	n, err := fmt.Sscanf(str, "bearer %s", &res)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("n: %d, res: %s\n", n, res)
}

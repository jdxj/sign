package evo

import (
	"fmt"
	"testing"
	"time"
)

func TestBiString(t *testing.T) {
	bi := &buildInfo{
		Filename: "abc",
		Datetime: time.Now().Unix(),
		Size:     2014810956,
		URL:      "def",
	}
	fmt.Printf("%s\n", bi)
}

func TestUpdater_Execute(t *testing.T) {
	updater := &Updater{}
	msg, err := updater.Execute("raphael")
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", msg)
}

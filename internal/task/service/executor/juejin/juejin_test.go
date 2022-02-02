package juejin

import (
	"fmt"
	"testing"
)

func TestSignIn_Execute(t *testing.T) {
	si := &SignIn{}
	res, err := si.Execute(nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Println(res)
}

func TestCount_Execute(t *testing.T) {
	count := &Count{}
	res, err := count.Execute(nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Println(res)
}

func TestPoint_Execute(t *testing.T) {
	p := &Point{}
	res, err := p.Execute(nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Println(res)
}

func TestCalendar_Execute(t *testing.T) {
	cal := &Calendar{}
	res, err := cal.Execute(nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Println(res)
}

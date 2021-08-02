package storage

import (
	"fmt"
	"testing"
)

func TestStorage_Write(t *testing.T) {
	s := &Storage{
		path: "./tasks.bak",
	}
	err := s.Write([]byte("456def"))
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

func TestStorage_Read(t *testing.T) {
	s := &Storage{
		path: "./tasks.bak",
	}
	data, err := s.Read()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", data)
}

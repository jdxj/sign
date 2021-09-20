package task

import (
	"fmt"
	"net/http"
	"testing"
)

func TestCookieToString(t *testing.T) {
	req, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	_ = req
}

func TestManager_Marshal(t *testing.T) {
	task := &Task{
		ID:     "jdxj",
		Domain: 301,
		Types:  []int{302},
		Key:    "",
	}
	m := NewManager()
	err := m.Add(task)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	data, err := m.Marshal()
	if err != nil {
		t.Fatalf("%s\n", err)
	}
	fmt.Printf("%s\n", data)
	err = m.Unmarshal(data)
	if err != nil {
		t.Fatalf("%s\n", err)
	}
}

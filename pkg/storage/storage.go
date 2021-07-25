package storage

import (
	"net/http/cookiejar"

	"github.com/jdxj/sign/pkg/toucher"
)

var (
	DefaultStorage = NewMemory()
)

type UserData struct {
	ID     string
	Domain string

	jar *cookiejar.Jar
	toucher.Toucher
}

func NewMemory() *Memory {
	m := &Memory{
		uds: make(map[int]toucher.Toucher),
	}
	return m
}

type Memory struct {
	num int
	uds map[int]toucher.Toucher
}

func (m *Memory) AddToucher(tch toucher.Toucher) error {
	m.num++
	m.uds[m.num] = tch
	return nil
}

func (m *Memory) GetAllUserData() map[int]toucher.Toucher {
	return m.uds
}

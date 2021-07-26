package storage

import (
	"github.com/jdxj/sign/pkg/toucher"
)

var (
	Default = NewMemory()
)

func NewMemory() *Memory {
	m := &Memory{
		uds: make(map[int]toucher.Validator),
	}
	return m
}

type Memory struct {
	num int
	uds map[int]toucher.Validator
}

func (m *Memory) AddUserData(val toucher.Validator) {
	m.num++
	m.uds[m.num] = val
}

func (m *Memory) DelUserData(num int) {
	delete(m.uds, num)
}

func (m *Memory) GetAllUserData() map[int]toucher.Validator {
	return m.uds
}

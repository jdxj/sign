package storage

import (
	"net/http/cookiejar"

	"github.com/jdxj/sign/pkg/toucher"
)

var (
	Default = NewMemory()
)

func NewUserData(id, domain, key string) (*UserData, error) {
	var tch toucher.Toucher
	switch domain {
	case toucher.DomainBili:
		tch = toucher.BiliTch
	default:
		return nil, toucher.ErrorUnsupportedDomain
	}

	jar, err := tch.Auth(key)
	if err != nil {
		return nil, err
	}

	ud := &UserData{
		ID:     id,
		Domain: domain,
		jar:    jar,
		tch:    tch,
	}
	return ud, nil
}

type UserData struct {
	ID     string
	Domain string

	jar *cookiejar.Jar
	tch toucher.Toucher
}

func (ud *UserData) Execute() error {
	err := ud.tch.SignIn(ud.jar)
	if err != nil {
		return err
	}
	return ud.tch.Verify(ud.jar)
}

func NewMemory() *Memory {
	m := &Memory{
		uds: make(map[int]*UserData),
	}
	return m
}

type Memory struct {
	num int
	uds map[int]*UserData
}

func (m *Memory) AddUserData(ud *UserData) error {
	m.num++
	m.uds[m.num] = ud
	return nil
}

func (m *Memory) DelUserData(num int) {
	delete(m.uds, num)
}

func (m *Memory) GetAllUserData() map[int]*UserData {
	return m.uds
}

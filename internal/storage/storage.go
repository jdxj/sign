package storage

import (
	"io/ioutil"
	"os"
)

var (
	defaultStorage *Storage
)

func Init(path string) {
	defaultStorage = NewStorage(path)
}

func Write(data []byte) error {
	return defaultStorage.Write(data)
}

func Read() ([]byte, error) {
	return defaultStorage.Read()
}

func NewStorage(path string) *Storage {
	return &Storage{
		path: path,
	}
}

type Storage struct {
	path string
}

func (s *Storage) Write(data []byte) error {
	f, err := os.OpenFile(s.path, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Sync()
		_ = f.Close()
	}()

	_, err = f.Write(data)
	return err
}

func (s *Storage) Read() ([]byte, error) {
	f, err := os.OpenFile(s.path, os.O_RDONLY, 0666)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = f.Close()
	}()

	return ioutil.ReadAll(f)
}

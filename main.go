package main

import "errors"

// go run github.com/golang/mock/mockgen -source=main.go -destination=mocks/mock_main.go -package=mocks

//go:generate go run github.com/vektra/mockery/v2@v2.43.2 --name=Database

type Database interface {
	Get(key string) (string, error)
	Set(key, value string) error
}

type MockDatabase struct {
	Data map[string]string
}

func New() *MockDatabase {
	return &MockDatabase{
		Data: make(map[string]string),
	}
}

func (db *MockDatabase) Get(key string) (string, error) {
	value, ok := db.Data[key]
	if !ok {
		return "", errors.New("key not found")
	}
	return value, nil
}

func (db *MockDatabase) Set(key, value string) error {
	db.Data[key] = value
	return nil
}

package main

import (
	"errors"
	"fmt"
)

//go:generate mockgen -source=main.go -destination=mocks/mocks.go

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

//========================================= new code=======================================\\

type MyServer struct {
	db Database
}

func (serv *MyServer) PrintKeyDB(key string) {
	if serv.db == nil {
		return
	}

	value, err := serv.db.Get(key)

	if err != nil {
		return
	}

	fmt.Println(value)
}

func main() {
	serv := &MyServer{db: New()}

	if serv.db == nil {
		return
	}

	serv.db.Set("testkey", "testvalue")
	serv.PrintKeyDB("testkey")
}

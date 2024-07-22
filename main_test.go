package main

import (
	mock_main "test3/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMyServer_PrintKeyDB(t *testing.T) {

	t.Run("PRINT TEST", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()
		dbmock := mock_main.NewMockDatabase(mockCtrl)
		serv := &MyServer{db: dbmock}

		dbmock.EXPECT().Get(gomock.Any()).Times(1)
		serv.PrintKeyDB("testkey1")

	})

}

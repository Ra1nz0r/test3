package main

import (
	"test3/mocks"
	"testing"
)

func TestMockDatabase_Get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		db      *MockDatabase
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test1",
			db:   New(),
			args: args{
				key: "Key",
			},
			want:    "Key1",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dataMock := mocks.NewDatabase(t)

			dataMock.On("Get", tt.args.key).Return(tt.want, nil)
			dataMock.Get(tt.args.key)

			tt.db.Set("Key", "Key1")

			got, err := tt.db.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("MockDatabase.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("MockDatabase.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMockDatabase_Set(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		db      *MockDatabase
		args    args
		wantErr bool
	}{
		{
			name: "test2",
			db:   New(),
			args: args{
				key:   "Key",
				value: "Set",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		dataMock := mocks.NewDatabase(t)

		dataMock.On("Set", tt.args.key, tt.args.value).Return(nil)
		dataMock.Set(tt.args.key, tt.args.value)

		t.Run(tt.name, func(t *testing.T) {
			if err := tt.db.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("MockDatabase.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

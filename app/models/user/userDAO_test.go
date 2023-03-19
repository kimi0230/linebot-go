package models

import (
	"linebot-go/services/mongodb"
	"testing"
)

var (
	IP       = "127.0.0.1"
	PORT     = "27017"
	USERNAME = "user"
	PASSWORD = "user_password"
	POOLSIZE = "100"
	DATABASE = "linebot-go"
	COLNAME  = "users"
)

func TestCreate(t *testing.T) {
	var tests = []struct {
		arg1 UserDTO
	}{
		{
			UserDTO{
				UserID:        "5566",
				DisplayName:   "KK",
				PictureURL:    "123456",
				StatusMessage: "status message",
				Language:      "en",
			},
		},
	}
	mgClient, err := mongodb.ConnectMongoDB(IP, PORT, USERNAME, PASSWORD, POOLSIZE, DATABASE)
	if err != nil {
		t.Errorf("ConnectMongoDB :%v", err)
	}

	userDAO := NewUserDAOwithName(mgClient, DATABASE, COLNAME)

	for _, tt := range tests {
		if _, err := userDAO.Create(&tt.arg1); err != nil {
			t.Errorf("%v", err)
		}
	}
}

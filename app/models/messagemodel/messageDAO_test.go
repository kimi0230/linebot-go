package messagemodel

import (
	"linebot-go/services/mongodb"
	"testing"
	"time"
)

var (
	IP       = "127.0.0.1"
	PORT     = "27017"
	USERNAME = "user"
	PASSWORD = "user_password"
	POOLSIZE = "100"
	DATABASE = "linebot-go"
	COLNAME  = "messages"
)

func TestCreate(t *testing.T) {
	var tests = []struct {
		arg1 MessageDTO
	}{
		{
			MessageDTO{
				Type:        "message",
				UserID:      "5566",
				ReplyToken:  "123456789",
				MessageID:   "message id",
				MessageText: "Hello World, Kimi",
				Timestamp:   time.Now(),
			},
		},
	}
	mgClient, err := mongodb.ConnectMongoDB(IP, PORT, USERNAME, PASSWORD, POOLSIZE, DATABASE)
	if err != nil {
		t.Errorf("ConnectMongoDB :%v", err)
	}

	messageDAO := NewMessageDAOwithName(mgClient, DATABASE, COLNAME)

	for _, tt := range tests {
		if _, err := messageDAO.Create(&tt.arg1); err != nil {
			t.Errorf("%v", err)
		}
	}
}

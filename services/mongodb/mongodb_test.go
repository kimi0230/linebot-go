package mongodb

import (
	"testing"
)

func TestConnectMongoDB(t *testing.T) {
	var tests = []struct {
		ip       string
		port     string
		username string
		password string
		poolsize string
		database string
	}{
		{
			"127.0.0.1",
			"27017",
			"root",
			"root",
			"100",
			"admin",
		}, {
			"127.0.0.1",
			"27017",
			"user",
			"user_password",
			"100",
			"linebot-go",
		},
	}
	for _, tt := range tests {
		if _, err := ConnectMongoDB(tt.ip, tt.port, tt.username, tt.password, tt.poolsize, tt.database); err != nil {
			t.Errorf("%v", err)
		}
	}
}

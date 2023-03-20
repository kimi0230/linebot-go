package linesdk

import (
	"log"
	"testing"
)

var (
	CHANNEL_SECRET = "your channel secret"
	ACCESS_TOKEN   = "your access token"
)

var bot LineBot

func init() {
	var err error
	bot, err = NewLineBot(CHANNEL_SECRET, ACCESS_TOKEN)
	if err != nil {
		log.Println(err)
	}
}

func TestPushMessage(t *testing.T) {
	var tests = []struct {
		userId  string
		message string
	}{
		{
			"user id",
			"Hello Kimi",
		},
	}
	for _, tt := range tests {
		if _, err := bot.PushMessage(tt.userId, tt.message); err != nil {
			t.Errorf("%v", err)
		}
	}
}

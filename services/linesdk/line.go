package linesdk

import (
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

type LineBot struct {
	Client *linebot.Client
}

func NewLineBot(secret, token string) (LineBot, error) {
	bot, err := linebot.New(secret, token)
	return LineBot{Client: bot}, err
}

func (lb *LineBot) HandleMessage() {

}

func (lb *LineBot) ReplyMessage() {

}

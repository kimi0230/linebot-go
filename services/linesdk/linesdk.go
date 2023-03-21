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

func (lb *LineBot) ReplyMessage(replyToken string, messages string) (*linebot.BasicResponse, error) {
	basicResponse, err := lb.Client.ReplyMessage(replyToken, linebot.NewTextMessage(messages)).Do()
	if err != nil {

		return nil, err
	}
	return basicResponse, nil
}

// https://developers.line.biz/en/reference/messaging-api/#get-profile
func (lb *LineBot) GetProfile(userID string) (*linebot.UserProfileResponse, error) {
	profile, err := lb.Client.GetProfile(userID).Do()
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (lb *LineBot) PushMessage(userID, message string) (*linebot.BasicResponse, error) {
	response, err := lb.Client.PushMessage(userID, linebot.NewTextMessage(message)).Do()
	if err != nil {
		return nil, err
	}
	return response, nil
}

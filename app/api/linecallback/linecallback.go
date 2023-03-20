package linecallback

import (
	"fmt"
	"linebot-go/app/models/messagemodel"
	"linebot-go/app/models/usermodel"
	"linebot-go/services/linesdk"
	"linebot-go/services/mongodb"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

func Callback(c *gin.Context) {
	bot, err := linesdk.NewLineBot(viper.GetString("line.channel_secret"), viper.GetString("line.accsss_token"))
	if err != nil {
		log.Fatal(err)
	}

	events, err := bot.Client.ParseRequest(c.Request)

	if err != nil {
		log.Println(err)
		// Signature validation
		if err == linebot.ErrInvalidSignature {
			c.AbortWithStatus(http.StatusBadRequest)
		} else {
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		return
	}

	// Handle received events
	for _, event := range events {
		fmt.Printf("event = %+v \n", event)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// echo message
				userID := event.Source.UserID

				// Create mongo client
				mgClient, err := mongodb.NewMongoClient()
				if err != nil {
					log.Print(err)
					c.AbortWithStatus(http.StatusInternalServerError)
				}

				// Get user profile and create or update
				profile, err := bot.GetProfile(userID)
				if err != nil {
					log.Print(err)
					c.AbortWithStatus(http.StatusInternalServerError)
				} else {
					displayName := profile.DisplayName
					photoURL := profile.PictureURL
					statusMessage := profile.StatusMessage
					fmt.Printf("User: name=%s\t, photo=%s\t, status=%s\n", displayName, photoURL, statusMessage)

					// Create or Update user info in mongodb
					userDAO := usermodel.NewUserDAO(mgClient)
					userDTO := usermodel.UserDTO{
						// ID:            primitive.NewObjectID(),
						UserID:        userID,
						DisplayName:   profile.DisplayName,
						PictureURL:    profile.PictureURL,
						StatusMessage: profile.StatusMessage,
						Language:      profile.Language,
					}
					userDAO.CreateOrUpdateByUserID(&userDTO)
				}

				// Create message info in mongodb
				messageDAO := messagemodel.NewMessageDAO(mgClient)
				messageDTO := messagemodel.MessageDTO{
					Type:        string(event.Type),
					UserID:      userID,
					MessageID:   message.ID,
					MessageType: string(message.Type()),
					MessageText: message.Text,
				}
				messageDAO.Create(&messageDTO)

				// Reply message
				if _, err := bot.ReplyMessage(event.ReplyToken, message.Text); err != nil {
					log.Print(err)
					c.AbortWithStatus(http.StatusInternalServerError)
				}
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

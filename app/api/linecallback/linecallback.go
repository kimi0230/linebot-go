package linecallback

import (
	"context"
	"fmt"
	models "linebot-go/app/models/user"
	"linebot-go/services/linesdk"
	"linebot-go/services/mongodb"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		fmt.Println("event", event)
		if event.Type == linebot.EventTypeMessage {
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				// echo message
				userID := event.Source.UserID
				fmt.Printf("userID = %s", userID)

				// get user profile
				profile, err := bot.GetProfile(userID)
				if err != nil {
					log.Print(err)
					c.AbortWithStatus(http.StatusInternalServerError)
				} else {
					displayName := profile.DisplayName
					photoURL := profile.PictureURL
					statusMessage := profile.StatusMessage
					fmt.Printf("User: name=%s\t, photo=%s\t, status=%s\n", displayName, photoURL, statusMessage)

					// TODO: Create or Update user info in mongodb
					mgClient, _ := mongodb.NewMongoClient()
					ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
					defer cancel()
					mgClient.Ping(ctx, nil)

					userDAO := models.NewUserDAO(mgClient)

					userDTO := models.UserDTO{
						ID:            primitive.NewObjectID(),
						UserID:        userID,
						DisplayName:   profile.DisplayName,
						PictureURL:    profile.PictureURL,
						StatusMessage: profile.StatusMessage,
						Language:      profile.Language,
					}
					userDAO.CreateUser(&userDTO)
				}

				// Reply message
				if _, err := bot.ReplyMessage(event.ReplyToken, message.Text); err != nil {
					log.Print(err)
					c.AbortWithStatus(http.StatusInternalServerError)
				}

				// TODO: Create message info in mongodb
			}
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

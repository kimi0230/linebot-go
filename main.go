package main

import (
	"flag"
	"fmt"
	"linebot-go/services/linesdk"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/spf13/viper"
)

var (
	opts options
)

type options struct {
	configFile string
}

func init() {
	flag.StringVar(&opts.configFile, "c", "./config.toml", "path of config file")
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments] <command> \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	viper.AutomaticEnv()
	viper.SetConfigFile(opts.configFile)
	viper.ReadInConfig()

	fmt.Fprintf(os.Stderr, "App Name : %s\n", viper.GetString("app.name"))

	// TODO: Gin Server
	r := gin.Default()
	r.UseH2C = true

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/callback", func(c *gin.Context) {
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
					}

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
	})

	port := viper.GetString("http.port")
	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  viper.GetDuration("http.read_timeout"),
		WriteTimeout: viper.GetDuration("http.write_timeout"),
	}

	log.Fatal(httpServer.ListenAndServe())
}

package routes

import (
	"linebot-go/app/api/heartbeat"
	"linebot-go/app/api/linecallback"
	"linebot-go/app/api/message"
	"linebot-go/app/api/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", heartbeat.Ping)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")

		v1.POST("/callback", linecallback.Callback)

		v1.GET("/users", user.GetUsers)

		v1.GET("/messages", message.GetMessages)

	}

}

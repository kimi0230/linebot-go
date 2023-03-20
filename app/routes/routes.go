package routes

import (
	"linebot-go/app/api/heartbeat"
	"linebot-go/app/api/lineapi"
	"linebot-go/app/api/message"
	"linebot-go/app/api/user"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", heartbeat.Ping)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			v1.GET("/users", user.GetUsers)
			v1.GET("/messages", message.GetMessages)

			line := v1.Group("/line")
			{
				line.POST("/callback", lineapi.Callback)
				line.POST("/message/push", lineapi.PushMessage)
			}
		}

	}
}

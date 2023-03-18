package routes

import (
	"linebot-go/app/api/heartbeat"
	"linebot-go/app/api/linecallback"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", heartbeat.Ping)

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")

		v1.POST("/callback", linecallback.Callback)

		admin := v1.Group("/admin")
		{
			admin.GET("/messages", func(c *gin.Context) {
				//...
			})

			admin.POST("/users", func(c *gin.Context) {
				//...
			})
		}
	}

}

package heartbeat

import (
	"linebot-go/services/mongodb"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Ping(c *gin.Context) {
	if _, err := mongodb.ConnectMongoDB(viper.GetString("mongo.ip"), viper.GetString("mongo.port"), viper.GetString("mongo.username"), viper.GetString("mongo.password"), viper.GetString("mongo.poolsize"), viper.GetString("mongo.database")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

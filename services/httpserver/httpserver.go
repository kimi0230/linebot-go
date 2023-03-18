package httpserver

import (
	"linebot-go/app/routes"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func StartGinServer(port int) {
	r := gin.Default()
	r.UseH2C = true

	routes.RegisterRoutes(r)

	httpServer := &http.Server{
		Addr:         ":" + strconv.Itoa(port),
		Handler:      r,
		ReadTimeout:  viper.GetDuration("http.read_timeout"),
		WriteTimeout: viper.GetDuration("http.write_timeout"),
	}

	log.Fatal(httpServer.ListenAndServe())
}

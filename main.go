package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

	fmt.Println(viper.GetDuration("http.read_timeout"))
	fmt.Println(viper.GetDuration("http.write_timeout"))
	port := viper.GetString("http.port")
	httpServer := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  viper.GetDuration("http.read_timeout"),
		WriteTimeout: viper.GetDuration("http.write_timeout"),
	}

	log.Fatal(httpServer.ListenAndServe())
}

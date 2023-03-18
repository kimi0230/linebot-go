package main

import (
	"fmt"
	"os"

	"linebot-go/cmd"

	"github.com/spf13/viper"
)

func main() {
	cmd.Execute()

	fmt.Fprintf(os.Stderr, "App Name : %s\n", viper.GetString("app.name"))

}

package cmd

import (
	"fmt"
	"linebot-go/services/httpserver"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	port    int
)

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "Prints 'hello, world'",
	Long:  `A simple command that prints 'hello, world' to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello, world")
	},
}

var serverCmd = &cobra.Command{
	Use:   "http",
	Short: "start http server",
	Long:  `start http server with port`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("http server : %s\n", viper.GetString("app.name"))
		httpserver.StartGinServer(port)
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.toml", "config file (default is $PWD/config.toml)")
	serverCmd.Flags().IntVar(&port, "port", 8080, "default is 8080")
	rootCmd.AddCommand(serverCmd)
}

func initConfig() {

	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	app     string
	version string
	state   string
	num     int
}

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("./")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	conf := &config{
		app:     viper.GetString("app"),
		version: viper.GetString("version"),
		state:   viper.GetString("state"),
		num:     viper.GetInt("num"),
	}
	fmt.Println(conf)
}

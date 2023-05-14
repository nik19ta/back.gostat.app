package main

import (
	"gostat/pkg/config"
	"gostat/server"
	"log"
)

func main() {
	conf := config.GetConfig()

	app := server.NewApp()

	if err := app.Run(conf.Port); err != nil {
		log.Fatalf("%s", err.Error())
	}
}

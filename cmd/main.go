package main

import (
	"log"

	"github.com/joho/godotenv"
	"guthub.com/gofmanaa/telegram-bot/cmd/app"
	"guthub.com/gofmanaa/telegram-bot/pkg/config"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

}

func main() {

	conf := config.Load()
	app.Run(conf)
}

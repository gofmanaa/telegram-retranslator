package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"guthub.com/gofmanaa/telegram-bot/cmd/app"
	"guthub.com/gofmanaa/telegram-bot/config"
	"log"
	"time"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	conf := config.New()
	app.Run(conf)
	//for t := range time.Tick(time.Hour*2) {
	//	fmt.Printf("Start app.Run() at %v", t)
	//	app.Run(conf)
	//}

}

func doEvery(d time.Duration, f func(...interface{})) {
	for t := range time.Tick(d) {
		fmt.Printf("Start app.Run() at %v", t)
		f()
	}
}

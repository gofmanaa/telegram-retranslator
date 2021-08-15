package main

import (
	"github.com/joho/godotenv"
	"guthub.com/gofmanaa/telegram-bot/cmd/telegram-bot/app"
	"guthub.com/gofmanaa/telegram-bot/config"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"runtime"
)

import _ "github.com/mkevac/debugcharts"

func init() {
	runtime.SetBlockProfileRate(0)
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {

	go func() {

		http.ListenAndServe(":8080", nil)

	}()

	logFile, err := os.OpenFile("log/log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Error open log file")
	}

	defer logFile.Close()
	logger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	conf := config.New(logger)

	app.Run(conf)

}

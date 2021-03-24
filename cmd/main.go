package main

import (
	"fmt"
	"guthub.com/gofmanaa/telegram-bot/cmd/app"
	"time"
)

func main() {
	app.Run()
	//doEvery(time.Minute*1, app.Run)
}

func doEvery(d time.Duration, f func()) {
	for t := range time.Tick(d) {
		fmt.Printf("Start app.Run() at %v", t)
		f()
	}
}

package app

import (
	"fmt"
	"io/ioutil"

	"guthub.com/gofmanaa/telegram-bot/pkg/config"
	"guthub.com/gofmanaa/telegram-bot/pkg/parser"
)

func Run(conf *config.Configuration) {
	fmt.Println("Start application.")
	defer fmt.Println("Stop application.")

	file, err := ioutil.ReadFile("tests/sg.json")
	if err != nil {
		fmt.Println(err)
	}

	media := parser.Scan(file)
	err = media.Save("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	//server.Run(conf, media)
	//bot.Run(conf, media)

}

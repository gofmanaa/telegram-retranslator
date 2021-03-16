package app

import (
	"fmt"
	"guthub.com/gofmanaa/telegram-bot/pkg/parser"
	"io/ioutil"
)

func Run() {
	fmt.Println("Start application.")
	defer fmt.Println("Stop application.")

	file, err := ioutil.ReadFile("tests/sg.json")
	if err != nil {
		fmt.Println(err)
	}

	parser.Read(file)

	//	bot.Run()
}

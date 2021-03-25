package bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"guthub.com/gofmanaa/telegram-bot/pkg/config"
	"guthub.com/gofmanaa/telegram-bot/pkg/store"
)

func Run(conf *config.Configuration, media *store.Media) {
	fmt.Println(conf.TelegramApiToken)
	bot, err := tgbotapi.NewBotAPI(conf.TelegramApiToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	fmt.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}

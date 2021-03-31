package bot

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"guthub.com/gofmanaa/telegram-bot/config"
	"guthub.com/gofmanaa/telegram-bot/pkg/redis_db"
	"log"
	"time"
)

func Run(ctx context.Context, rdx *redis.Client, conf *config.Config) {
	bot, err := tgbotapi.NewBotAPI(conf.TelegramApiKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	var done chan struct{}
	ticker := time.NewTicker(time.Minute * 10)

	go func() {
		for t := range ticker.C {
			fmt.Printf("Start bot.Run() at %v", t)
			keys := redis_db.KeysByStatus(ctx, rdx, 0)
			if len(keys) > 0 {
				for i := 0; i < 5; i++ {
					url := redis_db.GetByKey(ctx, rdx, keys[i])
					fmt.Println(url)
					SendPhoto(bot, conf.TelegramSpaceGettoChatId, url) //todo use NewMediaGroup
					redis_db.Publish(ctx, rdx, url)
				}
			} else {
				ticker.Stop()
				done <- struct{}{}
			}
		}
		done <- struct{}{}
	}()
	<-done

}

func SendPhoto(bot *tgbotapi.BotAPI, chatId int, url string) {
	msg := tgbotapi.NewPhotoShare(int64(chatId), url)
	_, err := bot.Send(msg)
	if err != nil {
		log.Fatalln("Error Telegram send msg: ", err)
	}
}

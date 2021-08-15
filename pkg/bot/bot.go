package bot

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"guthub.com/gofmanaa/telegram-bot/config"
	"guthub.com/gofmanaa/telegram-bot/pkg/redis_db"
	"time"
)

const PhotosPerMessage = 1
const RepostTime = 10

func Run(ctx context.Context, rdx *redis.Client, conf *config.Config) {
	log := conf.Log
	bot, err := tgbotapi.NewBotAPI(conf.TelegramAPIKey)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	var done chan struct{}
	ticker := time.NewTicker(time.Minute * RepostTime)

	go func() {
		for t := range ticker.C {
			log.Printf("Start bot.Run() at %v\n", t)
			keys := redis_db.KeysByStatus(ctx, rdx, 0)
			if len(keys) > 0 {
				storageID := redis_db.GetByKey(ctx, rdx, keys[0])
				listURL := redis_db.SMembers(ctx, rdx, storageID)
				log.Printf("Post ID: [%s]\n", storageID)
				log.Printf("Url count: %d\n", len(listURL))
				SendMediaGroup(bot, conf.TelegramSpaceGettoChatID, listURL)
				redis_db.RenameSet(ctx, rdx, storageID)
			} else {
				ticker.Stop()
				done <- struct{}{}
			}
		}
		done <- struct{}{}
	}()

	<-done
}

func SendMediaGroup(bot *tgbotapi.BotAPI, chatID int, urls []string) {
	var files []interface{}
	for _, url := range urls {
		files = append(files, tgbotapi.NewInputMediaPhoto(url))
		if len(files) >= PhotosPerMessage {
			msg := tgbotapi.NewMediaGroup(int64(chatID), files)
			_, _ = bot.Send(msg)
			time.Sleep(time.Millisecond * 500)
			files = []interface{}{}
		}
	}
}

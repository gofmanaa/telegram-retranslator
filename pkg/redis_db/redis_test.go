package redis_db

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"guthub.com/gofmanaa/telegram-bot/config"
	"log"
	"testing"
)

var ctx = context.Background()

func TestNotImgUrl(t *testing.T) {
	assert := assert.New(t)
	testData := []string{
		"test_https://i.imgur.com/dJJe2N7.gifv",
		"test_https://yo.be/qFdvwSyd_zs.jpg",
		"test_https://www.host/comments.png",
	}
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := config.New()
	rdx := InitRedis(ctx, conf)

	defer func() {
		for _, url := range testData {
			err := rdx.Del(ctx, GenerateKey(url, 8)).Err()
			if err != nil {
				log.Fatal("Redis set error: ", err)
			}
		}
	}()

	for _, url := range testData {
		Save(ctx, rdx, url, 9)
	}
	for _, url := range testData {
		key := "9|" + url
		redisUrl := GetByKey(ctx, rdx, key)
		assert.Equal(url, redisUrl, "they should be equal")
	}
	for _, url := range testData {
		err := rdx.Rename(ctx, GenerateKey(url, 9), GenerateKey(url, 8)).Err()
		if err != nil {
			log.Fatal("Redis set error: ", err)
		}
	}
	for _, url := range testData {
		key := "8|" + url
		redisUrl := GetByKey(ctx, rdx, key)
		assert.Equal(url, redisUrl, "they should be equal")
	}
}

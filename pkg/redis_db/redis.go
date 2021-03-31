package redis_db

import (
	"bytes"
	"context"
	"github.com/go-redis/redis/v8"
	"guthub.com/gofmanaa/telegram-bot/config"
	"log"
	"strconv"
)

func InitRedis(ctx context.Context, conf *config.Config) *redis.Client {
	rdx := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // replace to config
		Password: conf.RedisPass,   // no password set
		DB:       0,                // use default DB
	})
	_, err := rdx.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis error:", err)
	}

	return rdx
}

func Save(ctx context.Context, rdx *redis.Client, url string, status int) {
	err := rdx.Set(ctx, GenerateKey(url, status), url, 0).Err()
	if err != nil {
		log.Fatal("Redis set error: ", err)
	}
}

func GenerateKey(data string, status int) string {
	var b bytes.Buffer
	b.WriteString(strconv.Itoa(status))
	b.WriteString("|")
	b.WriteString(data)

	return b.String()
}

func KeysByStatus(ctx context.Context, rdx *redis.Client, status int) []string {
	searchKey := GenerateKey("*", status)
	srt, err := rdx.Keys(ctx, searchKey).Result()
	if err != nil {
		log.Fatal("Redis get error: ", err)
	}

	return srt
}

func GetByKey(ctx context.Context, rdx *redis.Client, key string) string {

	srt, err := rdx.Get(ctx, key).Result()
	if err != nil {
		log.Fatal("Redis get error: ", err)
	}

	return srt
}

func Publish(ctx context.Context, rdx *redis.Client, data string) {
	err := rdx.Rename(ctx, GenerateKey(data, 0), GenerateKey(data, 1)).Err()
	if err != nil {
		log.Fatal("Redis set error: ", err)
	}
}

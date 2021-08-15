package redis_db

import (
	"bytes"
	"context"
	"fmt"
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

func Set(ctx context.Context, rdx *redis.Client, id int, status int) {
	idKey := strconv.Itoa(id)
	err := rdx.Set(ctx, GenerateSKey(idKey, strconv.Itoa(status)), idKey, 0).Err()
	if err != nil {
		log.Fatal("Redis set error: ", err)
	}
}

func Get(ctx context.Context, rdx *redis.Client, id int, status int) string {
	idKey := strconv.Itoa(id)
	redisKey := GenerateSKey(idKey, strconv.Itoa(status))
	res, err := rdx.Get(ctx, redisKey).Result()
	if err == redis.Nil {
		fmt.Printf("Redis Get key [%s] does not exist\n", redisKey)
	} else if err != nil {
		log.Fatal("Redis set error: ", err)
	}

	return res
}

func HasIdStatus(ctx context.Context, rdx *redis.Client, id int, status string) bool {
	idKey := strconv.Itoa(id)
	redisKey := GenerateSKey(idKey, status)
	res := rdx.Keys(ctx, redisKey).Val()

	return len(res) != 0
}

func SAdd(ctx context.Context, rdx *redis.Client, id int, value string) {
	err := rdx.SAdd(ctx, strconv.Itoa(id), value).Err()
	if err != nil {
		log.Fatal("Redis SAdd error: ", err)
	}
}

func SMembers(ctx context.Context, rdx *redis.Client, key string) []string {
	res, err := rdx.SMembers(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("Redis SMembers key [%s] does not exist\n", key)
	} else if err != nil {
		log.Fatal("Redis set error: ", err)
	}
	return res
}

func KeysByStatus(ctx context.Context, rdx *redis.Client, status int) []string {
	searchKey := GenerateSKey("*", strconv.Itoa(status))
	srt, err := rdx.Keys(ctx, searchKey).Result()
	if err != nil {
		log.Fatal("Redis get error: ", err)
	}

	return srt
}

func GenerateSKey(key, status string) string {
	var b bytes.Buffer
	b.WriteString(status)
	b.WriteString("|")
	b.WriteString(key)

	return b.String()
}

func GenerateKey(id string, data string, status string) string {
	var b bytes.Buffer
	b.WriteString(status)
	b.WriteString("|")
	b.WriteString(id)
	b.WriteString("|")
	b.WriteString(data)

	return b.String()
}

func RenameSet(ctx context.Context, rdx *redis.Client, id string) {
	res, err := rdx.Rename(ctx, GenerateSKey(id, "0"), GenerateSKey(id, "1")).Result()
	if err == redis.Nil {
		fmt.Printf("Redis RenameSet key [%s] does not exist\n", id)
	} else if err != nil {
		log.Fatal("Redis get error: ", err)
	}
	fmt.Println("Redis RenameSet:", res)
}

func GetByKey(ctx context.Context, rdx *redis.Client, key string) string {
	srt, err := rdx.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("Redis [%s] does not exist\n", key)
	} else if err != nil {
		log.Fatal("Redis get error: ", err)
	}

	return srt
}

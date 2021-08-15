package main

import (
	"context"
	"github.com/joho/godotenv"
	"guthub.com/gofmanaa/telegram-bot/config"
	"guthub.com/gofmanaa/telegram-bot/pkg/redis_db"
	"html/template"
	"log"
	"net/http"
	"os"
)

type ViewData struct {
	Body []string
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	var ctx = context.Background()
	logFile, err := os.OpenFile("log/log_http.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("Error open log file")
	}

	defer logFile.Close()
	logger := log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	conf := config.New(logger)
	rdb := redis_db.InitRedis(ctx, conf)
	vd := ViewData{}
	keys := redis_db.KeysByStatus(ctx, rdb, 1)
	for _, key := range keys {
		vd.Body = append(vd.Body, redis_db.GetByKey(ctx, rdb, key))
	}

	http.HandleFunc("/", Page(func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("web/templates/index.html")
		_ = tmpl.Execute(w, vd)
	}))
	_ = http.ListenAndServe("localhost:8080", nil)
}

func Page(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}

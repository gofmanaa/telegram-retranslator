package server

import (
	"html/template"
	"log"
	"net/http"

	"guthub.com/gofmanaa/telegram-bot/pkg/config"
	"guthub.com/gofmanaa/telegram-bot/pkg/store"
)

type ViewData struct {
	Title string
	*store.Media
}

func Run(conf *config.Configuration, data *store.Media) {

	page := &ViewData{Title: "Index"}
	page.Media = data
	http.HandleFunc("/", indexHandler(page))
	http.ListenAndServe(":8090", nil)
}

func indexHandler(viewData *ViewData) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		templates := template.Must(template.ParseGlob("./web/*"))

		err := templates.ExecuteTemplate(w, "index.html", viewData)
		if err != nil {
			log.Fatal(err)
		}

	}
}

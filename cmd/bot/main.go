package main

import (
	"fmt"
	"net/http"

	"github.com/indrasaputra/tetesan-hujan/internal/config"
	"github.com/indrasaputra/tetesan-hujan/internal/raindrop"
	"github.com/indrasaputra/tetesan-hujan/internal/telegram"
	"github.com/indrasaputra/tetesan-hujan/usecase"
)

func main() {
	cfg, cerr := config.NewConfig(".env")
	checkError(cerr)

	api := raindrop.NewAPI(cfg.Raindrop.BaseURL, cfg.Raindrop.Token)
	creator := usecase.NewRaindropCreator(api)
	bot, berr := telegram.NewBot(cfg.Telegram.WebhookURL, cfg.Telegram.Token, cfg.Telegram.OwnerID, creator)
	checkError(berr)

	webhook, werr := bot.Run()
	checkError(werr)

	http.HandleFunc("/", webhook.ServeHTTP)
	fmt.Printf("Listening on port %s\n", cfg.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), http.DefaultServeMux)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

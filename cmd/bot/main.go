package main

import (
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

	bot.Run()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

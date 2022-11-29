package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/indrasaputra/tetesan-hujan/internal/config"
	"github.com/indrasaputra/tetesan-hujan/internal/raindrop"
	"github.com/indrasaputra/tetesan-hujan/internal/service"
	"github.com/indrasaputra/tetesan-hujan/internal/telegram"
)

const (
	timeout = 30 * time.Second
)

func main() {
	cfg, cerr := config.NewConfig(".env")
	checkError(cerr)

	api := raindrop.NewAPI(cfg.Raindrop.BaseURL, cfg.Raindrop.Token)
	creator := service.NewRaindropCreator(api)
	bot, berr := telegram.NewBot(cfg.Telegram.WebhookURL, cfg.Telegram.Token, cfg.Telegram.OwnerID, creator)
	checkError(berr)

	webhook := bot.Run()

	http.HandleFunc("/", webhook.ServeHTTP)
	http.HandleFunc("/healthz", healthzHandler)
	fmt.Printf("Listening on port %s\n", cfg.Port)

	server := &http.Server{
		Addr:        fmt.Sprintf(":%s", cfg.Port),
		Handler:     http.DefaultServeMux,
		ReadTimeout: timeout,
	}

	server.ListenAndServe()
}

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`OK`))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

package telegram

import (
	telebot "gopkg.in/tucnak/telebot.v2"
)

// Bot acts as Telegram Bot.
type Bot struct {
	*telebot.Bot
}

// NewBot creates an instance of Telegram Bot.
func NewBot(webhookURL, token string) (*Bot, error) {
	webhook := &telebot.Webhook{
		Endpoint: &telebot.WebhookEndpoint{
			PublicURL: webhookURL,
		},
	}

	setting := telebot.Settings{
		Token:  token,
		Poller: webhook,
	}

	bot, err := telebot.NewBot(setting)
	return &Bot{bot}, err
}

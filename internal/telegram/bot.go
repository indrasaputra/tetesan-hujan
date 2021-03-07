package telegram

import (
	"context"
	"strings"

	"github.com/indrasaputra/tetesan-hujan/entity"
	"github.com/indrasaputra/tetesan-hujan/usecase"
	telebot "gopkg.in/tucnak/telebot.v2"
)

// Bot acts as Telegram Bot.
type Bot struct {
	*telebot.Bot
	ownerID    int
	bookmarker usecase.CreateBookmark
}

// NewBot creates an instance of Telegram Bot.
func NewBot(webhookURL, token string, ownerID int, bookmarker usecase.CreateBookmark) (*Bot, error) {
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
	return &Bot{bot, ownerID, bookmarker}, err
}

// Run runs Telegram Bot.
func (b *Bot) Run() {
	b.setupMiddleware()
	b.setupTextHandler()
	b.Start()
}

func (b *Bot) setupTextHandler() {
	b.Handle(telebot.OnText, func(message *telebot.Message) {
		texts := strings.Split(message.Text, " ")
		if len(texts) != 2 {
			b.Send(message.Sender, "I only receive text containing collection name and URL")
			return
		}

		bookmark := &entity.Bookmark{CollectionName: texts[0], URL: texts[1]}
		if err := b.bookmarker.Create(context.Background(), bookmark); err != nil {
			b.Send(message.Sender, "Error on saving bookmark: %s", err.Error())
			return
		}
		b.Send(message.Sender, "Bookmark saved! Visit your raindrop application to see it")
	})
}

func (b *Bot) setupMiddleware() {
	midd := telebot.NewMiddlewarePoller(b.Poller, func(update *telebot.Update) bool {
		if update.Message.Sender.ID != b.ownerID {
			b.Send(update.Message.Sender, "You are not my master. I only serve my master")
			return false
		}
		return true
	})
	b.Poller = midd
}

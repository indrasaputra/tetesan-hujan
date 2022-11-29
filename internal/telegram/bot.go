package telegram

import (
	"context"
	"fmt"
	"strings"

	"github.com/indrasaputra/tetesan-hujan/entity"
	"github.com/indrasaputra/tetesan-hujan/internal/service"
	telebot "gopkg.in/tucnak/telebot.v2"
)

const (
	numberOfWord = 2
)

// Bot acts as Telegram Bot.
type Bot struct {
	*telebot.Bot
	ownerID    int64
	bookmarker service.CreateBookmark
	webhook    *telebot.Webhook
}

// NewBot creates an instance of Telegram Bot.
func NewBot(webhookURL, token string, ownerID int64, bookmarker service.CreateBookmark) (*Bot, error) {
	webhook := &telebot.Webhook{
		Endpoint: &telebot.WebhookEndpoint{
			PublicURL: webhookURL,
		},
	}

	setting := telebot.Settings{
		Token: token,
	}

	bot, err := telebot.NewBot(setting)
	if err != nil {
		return nil, err
	}
	bot.Poller = webhook
	return &Bot{bot, ownerID, bookmarker, webhook}, err
}

// Run runs Telegram Bot.
func (b *Bot) Run() *telebot.Webhook {
	b.setupMiddleware()
	b.setupTextHandler()
	go b.Start()
	return b.webhook
}

func (b *Bot) setupTextHandler() {
	b.Handle(telebot.OnText, func(message *telebot.Message) {
		texts := strings.Split(message.Text, " ")
		if len(texts) != numberOfWord {
			b.Reply(message, "I only receive text containing collection name and URL")
			return
		}

		bookmark := &entity.Bookmark{CollectionName: texts[1], URL: texts[0]}
		if err := b.bookmarker.Create(context.Background(), bookmark); err != nil {
			msg := fmt.Sprintf("Error on saving bookmark: %s", err.Error())
			b.Reply(message, msg)
			return
		}
		b.Reply(message, "Bookmark saved! Visit your raindrop application to see it")
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

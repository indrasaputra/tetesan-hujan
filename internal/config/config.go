package config

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

// Telegram holds configuration related to Telegram bot.
type Telegram struct {
	OwnerID    int64  `env:"TELEGRAM_OWNER_ID,required"`
	WebhookURL string `env:"TELEGRAM_WEBHOOK_URL,required"`
	Token      string `env:"TELEGRAM_TOKEN,required"`
}

// Raindrop holds configuration related to Raindrop API.
type Raindrop struct {
	BaseURL string `env:"RAINDROP_BASE_URL,required"`
	Token   string `env:"RAINDROP_TOKEN,required"`
}

// Config holds configuration for the project.
type Config struct {
	Port     string `env:"PORT,default=8080"`
	Telegram Telegram
	Raindrop Raindrop
}

// NewConfig creates an instance of Config.
// It needs the path of the env file to be used.
func NewConfig(env string) (*Config, error) {
	godotenv.Load(env)

	var config Config
	if err := envdecode.Decode(&config); err != nil {
		return nil, errors.Wrap(err, "[NewConfig] error decoding env")
	}

	return &config, nil
}

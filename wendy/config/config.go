package config

import (
	"github.com/subosito/gotenv"
	"os"
)

type Config struct {
	AppPort  string
	BotToken string
	URL      string
}

func GetConfig() Config {
	gotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("WENDY_PORT")
	}

	botToken := os.Getenv("WENDY_TOKEN")
	if botToken == "" {
		panic("bot token is empty")
	}

	url := os.Getenv("WENDY_URL")
	if url == "" {
		panic("url is empty")
	}

	return Config{
		AppPort:  port,
		BotToken: botToken,
		URL:      url,
	}
}

package config

import (
	"github.com/subosito/gotenv"
	"os"
)

type Config struct {
	AppPort  string
	BotToken string
	Platform string
}

func GetConfig() Config {
	gotenv.Load()

	platform := os.Getenv("WENDY_PLATFORM")
	if platform == "" {
		platform = "TELEGRAM"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("WENDY_PORT")
	}

	botToken := os.Getenv("WENDY_TOKEN")
	if botToken == "" {
		panic("bot token is empty")
	}

	return Config{
		AppPort:  port,
		BotToken: botToken,
		Platform: platform,
	}
}

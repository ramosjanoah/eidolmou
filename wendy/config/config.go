package config

import (
	"github.com/subosito/gotenv"
	"os"
)

type Config struct {
	AppPort  string
	BotToken string
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

	return Config{
		AppPort:  port,
		BotToken: botToken,
	}
}

package config

import (
	"github.com/subosito/gotenv"
	"os"
)

type Config struct {
	AppPort  string
	BotToken string
	Platform string
	Toggle   bool
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

	toggleParam := os.Getenv("WENDY_TOGGLE")
	var toggle bool
	if toggleParam == "" {
		panic("toggle is empty")
	} else {
		if toggleParam == "1" {
			toggle = true
		} else {
			toggle = false
		}
	}

	return Config{
		AppPort:  port,
		BotToken: botToken,
		Platform: platform,
		Toggle:   toggle,
	}
}

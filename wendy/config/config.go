package config

import (
	"github.com/subosito/gotenv"
	"os"
)

var (
	AppPort  string
	Platform string

	BotToken  string
	BotToggle bool
)

func init() {
	gotenv.Load()

	Platform = os.Getenv("WENDY_PLATFORM")
	if Platform == "" {
		Platform = "TELEGRAM"
	}

	AppPort = os.Getenv("PORT")
	if AppPort == "" {
		AppPort = os.Getenv("WENDY_PORT")
	}

	BotToken = os.Getenv("WENDY_TOKEN")
	if BotToken == "" {
		panic("bot token is empty")
	}

	toggleParam := os.Getenv("WENDY_TOGGLE")
	if toggleParam == "" {
		panic("toggle is empty")
	} else {
		if toggleParam == "1" {
			BotToggle = true
		} else {
			BotToggle = false
		}
	}
}

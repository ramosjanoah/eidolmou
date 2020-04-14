package config

import (
	"github.com/subosito/gotenv"
	"os"
	"strconv"
)

var (
	AppPort  string
	Platform string

	BotToken  string
	BotToggle bool

	CurrentDir string

	HeartbeatURL    string
	HeartbeatToggle bool

	AdminID int64
)

func init() {
	gotenv.Load()

	// Platform
	Platform = os.Getenv("WENDY_PLATFORM")
	if Platform == "" {
		Platform = "TELEGRAM"
	}

	// App port
	AppPort = os.Getenv("PORT")
	if AppPort == "" {
		AppPort = os.Getenv("WENDY_PORT")
	}

	// Bot Token
	BotToken = os.Getenv("WENDY_TOKEN")
	if BotToken == "" {
		panic("bot token is empty")
	}

	// Bot Toggle
	toggleParam := os.Getenv("WENDY_BOT_TOGGLE")
	if toggleParam == "" {
		panic("Bot toggle is empty")
	} else {
		if toggleParam == "1" {
			BotToggle = true
		} else {
			BotToggle = false
		}
	}

	// Current directory
	dir, err := os.Getwd()
	if err != nil {
		panic("error when getting current directory")
	}
	CurrentDir = dir

	// Heartbeat
	toggleParam = os.Getenv("WENDY_HEARTBEAT_TOGGLE")
	if toggleParam == "" {
		panic("Heartbeattoggle is empty")
	} else {
		if toggleParam == "1" {
			HeartbeatToggle = true
		} else {
			HeartbeatToggle = false
		}
	}

	// AdminID
	adminIDStr := os.Getenv("WENDY_ADMIN_ID")
	adminIDInt, _ := strconv.Atoi(adminIDStr)
	AdminID = int64(adminIDInt)

}

package connection

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/yanzay/tbot/v2"
	"net/http"
	"time"
)

const telegramAPIURL = "https://api.telegram.org"

var (
	TbotClient *tbot.Client
	Heartbeat  *http.Client
)

func init() {
	if config.Platform == "TELEGRAM" {
		httpClient := http.Client{
			Timeout: 10 * time.Second,
		}
		TbotClient = tbot.NewClient(config.BotToken, &httpClient, telegramAPIURL)
	}

	if config.HeartbeatToggle {
		Heartbeat = &http.Client{Timeout: 5 * time.Second}
	}
}

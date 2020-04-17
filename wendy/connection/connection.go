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
)

func init() {
	if config.Platform == "TELEGRAM" {
		httpClient := http.Client{
			Timeout: 10 * time.Second,
		}
		TbotClient = tbot.NewClient(config.BotToken, &httpClient, telegramAPIURL)
	}
}

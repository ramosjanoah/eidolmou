package repository

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/yanzay/tbot/v2"
	"net/http"
	"strconv"
	"time"
)

type ActionBot struct {
	BotClient *tbot.Client
}

var (
	actionBot = &ActionBot{}
	ApiURL    = "https://api.telegram.org"
)

func init() {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	botClient := tbot.NewClient(config.BotToken, &httpClient, ApiURL)

	actionBot.BotClient = botClient
}

func NewActionBot() *ActionBot {
	return actionBot
}

func (a *ActionBot) SendMessage(targetID int64, message string) error {

	_, err := a.BotClient.SendMessage(strconv.Itoa(int(targetID)), message)
	if err == nil {
		return err
	}

	return nil
}

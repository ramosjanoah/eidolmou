package telegrambot

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
	ApiURL = "https://api.telegram.org"
)

func NewActionBot() *ActionBot {

	actionBot := &ActionBot{}
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	botClient := tbot.NewClient(config.BotToken, &httpClient, ApiURL)
	actionBot.BotClient = botClient

	return actionBot
}

func (a *ActionBot) SendMessage(targetID int64, message string) error {

	_, err := a.BotClient.SendMessage(strconv.Itoa(int(targetID)), message)
	return err

}

func (a *ActionBot) SendAnimation(targetID int64, animationURL string) error {
	_, err := a.BotClient.SendAnimation(strconv.Itoa(int(targetID)), animationURL)
	return err
}

func (a *ActionBot) SendAnimationFile(targetID int64, animationFilename string) error {
	_, err := a.BotClient.SendAnimationFile(strconv.Itoa(int(targetID)), animationFilename)
	return err
}

package actionbot

import (
	"github.com/ramosjanoah/eidolmou/wendy/connection"
	"github.com/yanzay/tbot/v2"
	"strconv"
)

type TelegramActionBot struct {
	BotClient *tbot.Client
}

func GetTelegramActionBot() *TelegramActionBot {
	return &TelegramActionBot{connection.TbotClient}
}

func (a *TelegramActionBot) SendMessage(targetID int64, message string) error {
	_, err := a.BotClient.SendMessage(strconv.Itoa(int(targetID)), message)
	return err
}

func (a *TelegramActionBot) SendAnimation(targetID int64, animationURL string) error {
	_, err := a.BotClient.SendAnimation(strconv.Itoa(int(targetID)), animationURL)
	return err
}

func (a *TelegramActionBot) SendAnimationFile(targetID int64, animationFilename string) error {
	_, err := a.BotClient.SendAnimationFile(strconv.Itoa(int(targetID)), animationFilename)
	return err
}

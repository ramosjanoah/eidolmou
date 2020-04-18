package actionbot

import (
	"github.com/ramosjanoah/eidolmou/wendy/connection"
	"github.com/ramosjanoah/eidolmou/wendy/model/message"
	"github.com/yanzay/tbot/v2"
	"strconv"
)

type TelegramActionBot struct {
	BotClient *tbot.Client
}

func GetTelegramActionBot() *TelegramActionBot {
	return &TelegramActionBot{connection.TbotClient}
}

func (a *TelegramActionBot) SendString(targetID int64, str string) error {
	_, err := a.BotClient.SendMessage(strconv.Itoa(int(targetID)), str)
	return err
}
func (a *TelegramActionBot) SendMessage(targetID int64, m *message.Message) error {
	_, err := a.BotClient.SendMessage(strconv.Itoa(int(targetID)), m.GetString(), tbot.OptParseModeMarkdown)
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

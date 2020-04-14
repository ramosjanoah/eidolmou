package repository

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/repository/telegrambot"
)

type ActionBot interface {
	SendMessage(targetID int64, message string) error
	SendAnimation(targetID int64, animationURL string) error
	SendAnimationFile(targetID int64, animiationFileName string) error
}

func NewActionBot() ActionBot {
	if config.Platform == "TELEGRAM" {
		return telegrambot.NewActionBot()
	}

	panic("Platform value is nil")
}

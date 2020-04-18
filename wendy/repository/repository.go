package repository

import (
	"github.com/ramosjanoah/eidolmou/wendy/model/message"
	"github.com/ramosjanoah/eidolmou/wendy/model/tgif"
)

type ActionBotRepository interface {
	SendString(targetID int64, message string) error
	SendMessage(targetID int64, txt *message.Message) error
	SendAnimation(targetID int64, animationURL string) error
	SendAnimationFile(targetID int64, animiationFileName string) error
}

type GifRepository interface {
	Check() error
	Insert(params tgif.CreateParams) (*tgif.TGif, error)
}

package repository

import (
	"github.com/ramosjanoah/eidolmou/wendy/model/tgif"
)

type ActionBotRepository interface {
	SendMessage(targetID int64, message string) error
	SendAnimation(targetID int64, animationURL string) error
	SendAnimationFile(targetID int64, animiationFileName string) error
}

type GifRepository interface {
	Check() error
	Insert(params tgif.CreateParams) (*tgif.TGif, error)
}

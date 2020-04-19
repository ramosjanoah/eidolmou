package repository

import (
	"github.com/ramosjanoah/eidolmou/wendy/model/message"
	"github.com/ramosjanoah/eidolmou/wendy/model/tgif"
	"github.com/ramosjanoah/eidolmou/wendy/sctx"
)

type ActionBotRepository interface {
	SendString(sctx sctx.Context, targetID int64, message string) error
	SendMessage(sctx sctx.Context, targetID int64, txt *message.Message) error
	SendAnimation(sctx sctx.Context, targetID int64, animationURL string) error
	SendAnimationFile(sctx sctx.Context, targetID int64, animiationFileName string) error
}

type GifRepository interface {
	Insert(sctx sctx.Context, params tgif.CreateParams) (*tgif.TGif, error)
	GetRandom(sctx sctx.Context) (*tgif.TGif, error)
}

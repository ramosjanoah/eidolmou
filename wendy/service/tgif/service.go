package tgif

import (
	"github.com/ramosjanoah/eidolmou/wendy/repository"
	"github.com/ramosjanoah/eidolmou/wendy/repository/actionbot"
	"github.com/ramosjanoah/eidolmou/wendy/repository/tgif"
)

// list all the repository needed for this service
var (
	ActionBotRepository repository.ActionBotRepository
	TGifRepository      repository.GifRepository
)

// initialize the repository
func init() {
	ActionBotRepository = actionbot.GetActionBot()
	TGifRepository = tgif.GetGifRepository()
}

type TGif struct {
	Name          string
	FileID        string
	AdderUsername string
	Valid         string
}

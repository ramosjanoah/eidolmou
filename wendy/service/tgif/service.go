package tgif

import (
	tgifModel "github.com/ramosjanoah/eidolmou/wendy/model/tgif"
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
	ID            int64
	Name          string
	FileID        string
	AdderUsername string
	Valid         bool
}

func parseFromModel(modelObject *tgifModel.TGif) TGif {
	return TGif{
		ID:     modelObject.ID,
		Name:   modelObject.Name,
		FileID: modelObject.Name,
		Valid:  modelObject.Valid,
	}
}

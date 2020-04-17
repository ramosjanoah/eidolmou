package tgif

import (
	"github.com/ramosjanoah/eidolmou/wendy/repository"
	"github.com/ramosjanoah/eidolmou/wendy/repository/actionbot"
)

// list all the repository needed for this service
var ActionBot repository.ActionBotRepository

// initialize the repository
func init() {
	ActionBot = actionbot.GetActionBot()
}

type TGif struct {
	Name          string
	FileID        string
	AdderUsername string
	Valid         string
}

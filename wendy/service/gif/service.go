package gifservice

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

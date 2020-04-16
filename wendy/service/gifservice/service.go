package gifservice

import (
	"github.com/ramosjanoah/eidolmou/wendy/repository"
)

// list all the repository needed for this service
var ActionBotRepository repository.ActionBotRepository

// initialize the repository
func init() {
	ActionBotRepository := repository.GetActionBot()
}

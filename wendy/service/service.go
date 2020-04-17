package service

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/repository"
	"github.com/ramosjanoah/eidolmou/wendy/repository/actionbot"
)

// list all the repository needed for this service
var ActionBot repository.ActionBotRepository

// initialize the repository
func init() {
	ActionBot = actionbot.GetActionBot()
}

func AreYouOK(targetID int64) (string, error) {
	// send message response for 'are you ok'
	err := ActionBot.SendMessage(targetID, AreYouOKResponseMsg)
	if err != nil {
		return "", err
	}

	return AreYouOKResponseMsg, nil
}

func PleaseCheckMe() error {
	if config.AdminID == 0 {
		return nil
	}

	// send check me message to admin
	return ActionBot.SendMessage(config.AdminID, PleaseCheckMeMsg)
}

func ErrorResponseCallback(targetID int64, err error) error {
	return ActionBot.SendMessage(targetID, err.Error())
}

func sugarCoatError(err error) string {
	return fmt.Sprintf("Damn :(\nyou got some error: '%s'", err.Error())
}

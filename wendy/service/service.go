package service

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/repository"
	"github.com/ramosjanoah/eidolmou/wendy/repository/actionbot"
	"math/rand"
	"time"
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

func SendMeGif(targetID int64) error {
	// pick random animation file
	rand.Seed(time.Now().Unix())
	pickedFile := fmt.Sprintf("%s/wendy/asset/%s.gif", config.CurrentDir, WendyGifs[rand.Int()%len(WendyGifs)])
	return ActionBot.SendAnimationFile(targetID, pickedFile)
}

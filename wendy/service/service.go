package service

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/repository"
	"math/rand"
	"time"
)

// list all the repository needed for this service
var ActionBotRepository repository.ActionBotRepository

// initialize the repository
func init() {
	ActionBotRepository := repository.GetActionBot()
}

func AreYouOK(targetID int64) (string, error) {
	// send message response for 'are you ok'
	err := ActionBotRepository.SendMessage(targetID, AreYouOKResponseMsg)
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
	return ActionBotRepository.SendMessage(config.AdminID, PleaseCheckMeMsg)
}

func SendMeGif(targetID int64) error {
	// pick random animation file
	rand.Seed(time.Now().Unix())
	pickedFile := fmt.Sprintf("%s/wendy/asset/%s.gif", config.CurrentDir, WendyGifs[rand.Int()%len(WendyGifs)])
	return ActionBotRepository.SendAnimationFile(targetID, pickedFile)

}

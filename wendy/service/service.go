package service

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/repository"
	"math/rand"
	"time"
)

type Service struct {
	ActionBot repository.ActionBot
}

var (
	service = Service{}

	WendyGifs = []string{
		"wendy-hi",
		"wendy-swag",
		"wendy-love-thumbs",
	}

	AreYouOKResponseMsg = "Hi! I'm ok, don't worry :)"
	PleaseCheckMeMsg    = "Hi, check me please :("
)

func init() {
	service.ActionBot = repository.NewActionBot()
}

func AreYouOK(targetID int64) (string, error) {
	// send message response for 'are you ok'
	err := service.ActionBot.SendMessage(targetID, AreYouOKResponseMsg)
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
	return service.ActionBot.SendMessage(config.AdminID, PleaseCheckMeMsg)
}

func SendMeGif(targetID int64) error {
	// pick random animation file
	rand.Seed(time.Now().Unix())
	pickedFile := fmt.Sprintf("%s/wendy/asset/%s.gif", config.CurrentDir, WendyGifs[rand.Int()%len(WendyGifs)])
	return service.ActionBot.SendAnimationFile(targetID, pickedFile)

}

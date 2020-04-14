package service

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/repository"
)

type Service struct {
	ActionBot repository.ActionBot
}

var (
	service = Service{}

	AreYouOKResponse = "Hi! I'm ok, don't worry :)"
)

func init() {
	service.ActionBot = repository.NewActionBot()
}

func AreYouOK(targetID int64) (string, error) {
	// There are no logic in AreYouOK so I'll just return the response

	err := service.ActionBot.SendMessage(targetID, AreYouOKResponse)
	if err != nil {
		return "", err
	}

	err = service.ActionBot.SendAnimationFile(targetID, config.CurrentDir+"/wendy/asset/wendy-hi.gif")
	if err != nil {
		return "", err
	}

	return AreYouOKResponse, nil
}

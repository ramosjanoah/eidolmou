package service

import (
	"github.com/ramosjanoah/eidolmou/wendy/repository"
)

type Service struct {
	ActionBot *repository.ActionBot
}

var (
	service          = Service{}
	AreYouOKResponse = "Hi! I'm ok, don't worry :)"
)

func init() {
	service.ActionBot = repository.NewActionBot()
}

func AreYouOK(targetID int64) (string, error) {
	// There are no logic in AreYouOK so I'll just return the response

	service.ActionBot.SendMessage(targetID, AreYouOKResponse)
	return AreYouOKResponse, nil
}

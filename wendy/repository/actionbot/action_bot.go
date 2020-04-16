package actionbot

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/repository"
)

func GetActionBot() repository.ActionBotRepository {
	if config.Platform == "TELEGRAM" {
		return GetTelegramActionBot()
	}

	panic("Platform value is nil")
}

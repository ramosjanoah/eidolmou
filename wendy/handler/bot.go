package handler

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
)

type Bot interface {
	ChatListen()
	HttpListen()
}

func NewBot() Bot {
	config := config.GetConfig()
	if config.Platform == "TELEGRAM" {
		return NewTelegramBot()
	} else {
		panic("Var 'plaftorm' not defined")
	}
}

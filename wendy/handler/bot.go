package handler

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/handler/telegram"
)

type Bot interface {
	ChatListen()
	HttpListen()
}

func NewBot() Bot {
	if config.Platform == "TELEGRAM" {
		return telegram.NewBot()
	} else {
		panic("Var 'plaftorm' not defined")
	}
}

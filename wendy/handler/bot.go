package handler

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/handler/telegram"
)

type Bot interface {
	ChatListen()
	HttpListen()
}

func NewHandlerBot() Bot {
	if config.Platform == "TELEGRAM" {
		return telegram.NewHandlerBot()
	} else {
		panic("Var 'plaftorm' not defined")
	}
}

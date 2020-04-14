package handler

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/handler/telegram"
)

type HandlerBot interface {
	ChatListen()
	HttpListen()
	Heartbeat()
}

func NewHandlerBot() HandlerBot {
	if config.Platform == "TELEGRAM" {
		return telegram.NewHandlerBot()
	} else {
		panic("Var 'plaftorm' not defined")
	}
}

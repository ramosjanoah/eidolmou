package handler

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/handler/http"
	"github.com/ramosjanoah/eidolmou/wendy/handler/telegram"
)

type Handler struct {
	chat ChatHandler
	http HttpHandler
}

type ChatHandler interface {
	ChatListen()
}

type HttpHandler interface {
	HttpListen()
}

func NewHandler() *Handler {
	httpHandler := http.NewHttpHandler()
	var chatHandler ChatHandler

	if config.Platform == "TELEGRAM" {
		chatHandler = telegram.NewChatHandler()
	} else {
		panic("Var 'plaftorm' not defined")
	}

	handler := Handler{
		http: httpHandler,
		chat: chatHandler,
	}

	return &handler
}

func (h *Handler) ChatListen() {
	h.chat.ChatListen()
}

func (h *Handler) HttpListen() {
	h.http.HttpListen()
}

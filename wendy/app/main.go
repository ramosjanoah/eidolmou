package main

import (
	"github.com/ramosjanoah/eidolmou/wendy/handler"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	bot := handler.NewBot()

	go bot.ChatListen()
	go bot.HttpListen()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-signals
}

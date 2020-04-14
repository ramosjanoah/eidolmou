package main

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/handler"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if config.BotToggle {
		bot := handler.NewBot()

		go bot.ChatListen()
		go bot.HttpListen()

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-signals
	} else {
		log.Println("Wendy toggled off")
	}
}

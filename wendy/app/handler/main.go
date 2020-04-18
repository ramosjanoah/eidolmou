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
	handler := handler.NewHandler()

	if config.BotToggle {
		go handler.ChatListen()
	} else {
		log.Println("Wendy's bot is toggled off.")
	}

	go handler.HttpListen()

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-signals
}

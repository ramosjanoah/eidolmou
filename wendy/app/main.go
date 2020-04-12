package main

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/handler"
	"log"
	// "net/http"
)

func main() {
	bot := handler.NewBot()
	config := config.GetConfig()
	log.Println(fmt.Sprintf("Wendy is listening in %s..", config.AppPort))
	bot.Start()
}

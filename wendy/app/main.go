package main

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/handler"
	"log"
	"net/http"
)

func main() {
	router := handler.NewRouter()
	config := config.GetConfig()
	log.Println(fmt.Sprintf("Wendy is listening in %s..", config.AppPort))
	http.ListenAndServe(":"+config.AppPort, router)
}

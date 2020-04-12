package main

import (
	"github.com/ramosjanoah/eidolmou/wendy/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	router := handler.NewRouter()
	log.Println("Wendy is listening..")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}

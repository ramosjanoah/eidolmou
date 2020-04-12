package main

import (
	"github.com/ramosjanoah/eidolmou/wendy/handler"
	"log"
	"net/http"
)

func main() {
	router := handler.NewRouter()
	log.Println("Wendy is listening in :80")
	http.ListenAndServe(":"+os.Getenv("PORT"), router)
}

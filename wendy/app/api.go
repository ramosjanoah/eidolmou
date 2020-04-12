package main

import (
	"github.com/ramosjanoah/eidolmou/wendy/handler"
	"log"
	"net/http"
)

func main() {
	router := handler.NewRouter()
	log.Println("Wendy is listening in :2345")
	http.ListenAndServe(":2345", router)
}

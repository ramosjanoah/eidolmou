package main

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/connection"
	"log"
	"net/http"
	"time"
)

var heartbeatURL = "https://white-wendy.herokuapp.com/wendy/areyouok"

func main() {
	if !config.HeartbeatToggle {
		log.Println("Heartbeat turned off")
		return
	}

	req, err := http.NewRequest("GET", heartbeatURL, nil)
	if err != nil {
		panic("Heartbeat failed")
	}

	for {
		_, err = connection.Heartbeat.Do(req)
		if err != nil {
			panic(fmt.Sprintf("Heartbeat failed, %s", err.Error()))
		}
		log.Println("Wendy's ok")
		time.Sleep(3 * time.Minute)
	}
}

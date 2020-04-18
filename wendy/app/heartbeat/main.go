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
	if !config.HeartbeatToggle {
		return
	}

	req, err := netHttp.NewRequest("GET", t.heartbeatURL, nil)
	if err != nil {
		panic("Heartbeat failed")
	}

	for {
		_, err = connection.Heartbeat.Do(req)
		if err != nil {
			panic(fmt.Sprintf("Heartbeat failed, %s", err.Error()))
		}
		time.Sleep(3 * time.Minute)
	}
}

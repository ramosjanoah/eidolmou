package telegram

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"log"
	"time"
)

func decorate(command string, handleFunc func(m *tbot.Message) error) (string, func(m *tbot.Message)) {

	decoratedFunction := func(m *tbot.Message) {
		startTime := time.Now()

		err := handleFunc(m)

		elapsedTime := time.Since(startTime).Seconds()
		if err != nil {
			log.Println(fmt.Sprintf(`{"state":"failed", "duration":%g, "caller":"%s"}`, elapsedTime, command))
		} else {
			log.Println(fmt.Sprintf(`{"state":"success", "duration":%g, "caller":"%s"}`, elapsedTime, command))
		}
	}

	return command, decoratedFunction
}

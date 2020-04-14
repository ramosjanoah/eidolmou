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
			log.Println(fmt.Sprintf(`{"state":"failed", "error": "%s", "duration":%g, "function_caller":"%s"}`, err.Error(), elapsedTime, command))
		} else {
			log.Println(fmt.Sprintf(`{"state":"success", "duration":%g, "function_caller":"%s"}`, elapsedTime, command))
		}
	}

	return command, decoratedFunction
}

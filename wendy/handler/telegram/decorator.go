package telegram

import (
	"fmt"
	"github.com/yanzay/tbot/v2"
	"log"
	"time"
)

func decorate(handleFunc func(m *tbot.Message) error) func(m *tbot.Message) {

	decoratedFunction := func(m *tbot.Message) {
		startTime := time.Now()

		err := handleFunc(m)

		elapsedTime := time.Since(startTime).Seconds()
		if err != nil {
			log.Println(fmt.Sprintf(`{"state":"success", "duration":%g}`, elapsedTime))
		} else {
			log.Println(fmt.Sprintf(`{"state":"success", "duration":%g}`, elapsedTime))
		}
	}

	return decoratedFunction
}

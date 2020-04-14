package telegram

import (
	"errors"
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/service"
	"github.com/yanzay/tbot/v2"
	"log"
	"time"
)

func decorate(command string, handleFunc func(m *tbot.Message) error) (string, func(m *tbot.Message)) {

	decoratedFunction := func(m *tbot.Message) {

		var err error
		startTime := time.Now()

		defer func() {
			elapsedTime := time.Since(startTime).Seconds()

			if r := recover(); r != nil {
				_ = service.PleaseCheckMe()
				err = errors.New(r.(string))
			}

			if err != nil {
				log.Println(fmt.Sprintf(`{"state":"failed", "msg": "%s", "duration":%g, "function_caller":"%s"}`, err.Error(), elapsedTime, command))
			} else {
				log.Println(fmt.Sprintf(`{"state":"success", "duration":%g, "function_caller":"%s"}`, elapsedTime, command))
			}
		}()

		// do the handling
		err = handleFunc(m)
	}

	return command, decoratedFunction
}

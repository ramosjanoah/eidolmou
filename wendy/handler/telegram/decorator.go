package telegram

import (
	"errors"
	"github.com/ramosjanoah/eidolmou/wendy/sctx"
	"github.com/ramosjanoah/eidolmou/wendy/service"
	"github.com/yanzay/tbot/v2"
	"time"
)

func decorate(command string, handleFunc func(ctx sctx.Context, m *tbot.Message) error) (string, func(m *tbot.Message)) {

	decoratedFunction := func(m *tbot.Message) {
		var err error

		sctx := sctx.New().WithTimeout(10 * time.Second)
		defer func() {
			if r := recover(); r != nil {
				err = errors.New(r.(string))
			}
			if err != nil {
				sctx.SetFinalError(err)
				_ = service.ErrorResponseCallback(sctx, int64(m.From.ID), err)
			}

			sctx.End()
			println(sctx.GetLogs())
		}()

		// do the handling
		err = handleFunc(sctx, m)
	}

	return command, decoratedFunction
}

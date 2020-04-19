package tgif

import (
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	"github.com/ramosjanoah/eidolmou/wendy/sctx"
)

func SendMeGif(sctx sctx.Context, targetID int64) (err error) {
	select {
	case <-sctx.Done():
		return errors.TimeoutError()
	default:
	}

	ctxl := sctx.AddLog("tgif Service", "SendMeGif").WithInfo("parameter:targetID", targetID)
	defer func() {
		ctxl.EndWithError(err)
	}()

	randomTGif, err := TGifRepository.GetRandom(sctx)
	if err != nil {
		return err
	}

	ctxl.WithInfo("random_t_gif", randomTGif)
	return ActionBotRepository.SendAnimation(sctx, targetID, randomTGif.FileID)
}

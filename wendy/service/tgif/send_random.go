package tgif

import (
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	"github.com/ramosjanoah/eidolmou/wendy/sctx"
)

func SendMeGif(sctx sctx.Context, targetID int64) error {
	select {
	case <-sctx.Done():
		return errors.TimeoutError()
	default:
	}

	randomTGif, err := TGifRepository.GetRandom(sctx)
	if err != nil {
		return err
	}

	return ActionBotRepository.SendAnimation(sctx, targetID, randomTGif.FileID)
}

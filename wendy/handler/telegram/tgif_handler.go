package telegram

import (
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	"github.com/ramosjanoah/eidolmou/wendy/sctx"
	tgifService "github.com/ramosjanoah/eidolmou/wendy/service/tgif"
	"github.com/yanzay/tbot/v2"
)

func (t *ChatHandler) sendMeGif(sctx sctx.Context, m *tbot.Message) (err error) {
	return tgifService.SendMeGif(sctx, int64(m.From.ID))
}

func (t *ChatHandler) addGif(sctx sctx.Context, m *tbot.Message) (err error) {

	if m.Caption == "" {
		return errors.PayloadMissingError()
	}
	_, args := parseCaption(m.Caption)

	if m.Document == nil {
		return errors.MissingParameterError("document: animation")
	}

	if m.From == nil {
		return errors.MissingParameterError("user")
	}

	addForm := tgifService.AddForm{
		Name:    args["name"],
		FileID:  m.Document.FileID,
		AdderID: int64(m.From.ID),
	}

	_, err = tgifService.Add(sctx, addForm)
	return err
}

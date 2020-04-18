package telegram

import (
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	tgifService "github.com/ramosjanoah/eidolmou/wendy/service/tgif"
	"github.com/yanzay/tbot/v2"
)

func (t *ChatHandler) sendMeGif(m *tbot.Message) (err error) {
	return tgifService.SendMeGif(int64(m.From.ID))
}

func (t *ChatHandler) addGif(m *tbot.Message) (err error) {

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

	_, err = tgifService.Add(addForm)
	return err
}

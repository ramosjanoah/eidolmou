package telegram

import (
	"errors"
	tgifService "github.com/ramosjanoah/eidolmou/wendy/service/tgif"
	"github.com/yanzay/tbot/v2"
)

func (t *HandlerBot) sendMeGif(m *tbot.Message) (err error) {
	return tgifService.SendMeGif(int64(m.From.ID))
}

// TODO: Create file of errors pool for error that used here
func (t *HandlerBot) addGif(m *tbot.Message) (err error) {

	if m.Caption == "" {
		return errors.New("info is nil")
	}
	parsedInfo := parseCaption(m.Caption)

	if m.Document == nil {
		return errors.New("Animation is nil")
	}

	if m.From == nil {
		return errors.New("User is nil")
	}

	addForm := tgifService.AddForm{
		Name:    parsedInfo["name"],
		FileID:  m.Document.FileID,
		AdderID: int64(m.From.ID),
	}

	_, err = tgifService.AddNewGif(addForm)

	if err != nil {
		return err
	}

	return nil
}

func parseCaption(str string) map[string]string {
	// TODO: create the good parseCaption here
	return map[string]string{
		"name": "temp",
	}
}

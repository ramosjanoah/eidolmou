package tgif

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	"github.com/ramosjanoah/eidolmou/wendy/model/message"
	tgifModel "github.com/ramosjanoah/eidolmou/wendy/model/tgif"
	"github.com/ramosjanoah/eidolmou/wendy/sctx"
)

type AddForm struct {
	Name    string
	FileID  string
	AdderID int64
}

func (f *AddForm) Validate() (err error) {
	if f.Name == "" {
		return errors.MissingParameterError("name")
	}

	if f.FileID == "" {
		return errors.MissingParameterError("document: animation")
	}

	if f.AdderID == 0 {
		return errors.MissingParameterError("user")
	}

	return nil

}

func Add(sctx sctx.Context, form AddForm) (TGif, error) {
	select {
	case <-sctx.Done():
		return TGif{}, errors.TimeoutError()
	default:
	}

	err := form.Validate()
	if err != nil {
		return TGif{}, err
	}

	tgifModel, err := TGifRepository.Insert(sctx, tgifModel.CreateParams{
		Name:    form.Name,
		FileID:  form.FileID,
		AdderID: form.AdderID,
	})
	if err != nil {
		return TGif{}, err
	}

	tgif := parseFromModel(tgifModel)

	m := message.New("Congrats, you succeeded to add this gif :)").
		AddNewLine(fmt.Sprintf("*ID:* %d", tgif.ID)).
		AddNewLine(fmt.Sprintf("*Adder:* '%s'", tgif.Name))

	err = ActionBotRepository.SendMessage(sctx, form.AdderID, m)
	if err != nil {
		return TGif{}, err
	}

	return tgif, err
}

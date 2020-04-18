package tgif

import (
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	"github.com/ramosjanoah/eidolmou/wendy/model/tgif"
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

func Add(form AddForm) (TGif, error) {
	err := form.Validate()
	if err != nil {
		return TGif{}, err
	}

	_, err = TGifRepository.Insert(tgif.CreateParams{
		Name:    form.Name,
		FileID:  form.FileID,
		AdderID: form.AdderID,
	})
	if err != nil {
		return TGif{}, err
	}

	err = ActionBotRepository.SendMessage(form.AdderID, "Congrats, you succeeded to add this gif :)")
	return TGif{}, err
}

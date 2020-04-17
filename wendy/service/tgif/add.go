package tgif

import (
	"github.com/ramosjanoah/eidolmou/wendy/errors"
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

func AddNewGif(form AddForm) (TGif, error) {
	err := form.Validate()
	if err != nil {
		return TGif{}, err
	}

	ActionBot.SendMessage(form.AdderID, "Congrats, you succeeded to add this gif :)")
	return TGif{}, nil
}

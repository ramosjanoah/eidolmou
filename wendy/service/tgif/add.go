package tgif

import (
	"log"
)

type AddForm struct {
	Name    string
	FileID  string
	AdderID int64
}

// TODO: create form validation here

func AddNewGif(form AddForm) (TGif, error) {
	// TODO: add gif to psql here
	log.Println("Congrats, you got here :)")
	log.Println(form)
	return TGif{}, nil
}

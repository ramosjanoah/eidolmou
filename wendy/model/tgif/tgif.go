package tgif

import "time"

type TGif struct {
	ID                int64     `json:"id" gorm:"primary_key"`
	Name              string    `json:"name" gorm:"unique"`
	FileID            string    `json:"file_id" gorm:"unique"`
	AdderID           int64     `json:"adder_id"`
	Valid             bool      `json:"valid"`
	LastValidityCheck time.Time `json:"last_validity_check"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type CreateParams struct {
	Name    string `json:"name"`
	FileID  string `json:"file_id"`
	AdderID int64  `json:"adder_id"`
}

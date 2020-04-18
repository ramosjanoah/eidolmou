package tgif

import "time"

type TGif struct {
	ID                int64  `gorm:"primary_key"`
	Name              string `gorm:"unique"`
	FileID            string `gorm:"unique"`
	AdderID           int64
	Valid             bool
	LastValidityCheck time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type CreateParams struct {
	Name    string
	FileID  string
	AdderID int64
}

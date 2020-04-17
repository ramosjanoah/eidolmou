package tgif

import (
	"time"
)

type TGif struct {
	ID                int64
	Name              string
	FileID            string
	AdderID           int64
	Valid             bool
	LastValidityCheck time.Time
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

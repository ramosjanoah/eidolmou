package tgif

import (
	"github.com/jinzhu/gorm"
	"github.com/ramosjanoah/eidolmou/wendy/connection"
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	"github.com/ramosjanoah/eidolmou/wendy/model/tgif"
	"time"
)

type GifSqlImplmentation struct {
	conn *gorm.DB
}

func GetGifSqlRepository() *GifSqlImplmentation {
	return &GifSqlImplmentation{conn: connection.SQLConn}
}

func (g *GifSqlImplmentation) Check() error { return nil }

func (g *GifSqlImplmentation) Insert(params tgif.CreateParams) (*tgif.TGif, error) {
	res := g.conn.Create(&tgif.TGif{
		Name:              params.Name,
		FileID:            params.FileID,
		AdderID:           params.AdderID,
		Valid:             true,
		LastValidityCheck: time.Now(),
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	})

	if res.Error != nil {
		return &tgif.TGif{}, errors.ErrorWhenInsertRow("TGif", res.Error)
	}

	return res.Value.(*tgif.TGif), nil
}

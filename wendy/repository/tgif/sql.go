package tgif

import (
	"github.com/jinzhu/gorm"
	"github.com/ramosjanoah/eidolmou/wendy/connection"
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	"github.com/ramosjanoah/eidolmou/wendy/model/tgif"
	"github.com/ramosjanoah/eidolmou/wendy/sctx"
	"time"
)

var model = &tgif.TGif{}

type GifSqlImplmentation struct {
	conn *gorm.DB
}

func GetGifSqlRepository() *GifSqlImplmentation {
	return &GifSqlImplmentation{conn: connection.SQLConn}
}

func (g *GifSqlImplmentation) Insert(sctx sctx.Context, params tgif.CreateParams) (*tgif.TGif, error) {
	select {
	case <-sctx.Done():
		return &tgif.TGif{}, errors.TimeoutError()
	default:
	}

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

func (g *GifSqlImplmentation) GetRandom(sctx sctx.Context) (*tgif.TGif, error) {
	select {
	case <-sctx.Done():
		return &tgif.TGif{}, errors.TimeoutError()
	default:
	}

	randomRows := &[]tgif.TGif{}
	res := g.conn.Limit(1).Order(gorm.Expr("rand()")).Find(randomRows)

	if res.Error != nil {
		return nil, res.Error
	}

	if randomRows == nil || len(*randomRows) == 0 {
		return nil, errors.NotFoundError("TGif")
	}

	return &(*randomRows)[0], nil
}

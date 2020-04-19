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

func (g *GifSqlImplmentation) Insert(sctx sctx.Context, params tgif.CreateParams) (t *tgif.TGif, err error) {
	select {
	case <-sctx.Done():
		return &tgif.TGif{}, errors.TimeoutError()
	default:
	}

	ctxl := sctx.AddLog("GifSqlImplmentation", "Insert").
		WithInfo("parameter:params", params)
	defer func() {
		ctxl.EndWithError(err).WithInfo("tgif", t)
	}()

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

	t = res.Value.(*tgif.TGif)
	return t, nil
}

func (g *GifSqlImplmentation) GetRandom(sctx sctx.Context) (t *tgif.TGif, err error) {
	select {
	case <-sctx.Done():
		return &tgif.TGif{}, errors.TimeoutError()
	default:
	}

	ctxl := sctx.AddLog("GifSqlImplmentation", "GetRandom")
	defer func() {
		ctxl.EndWithError(err).WithInfo("tgif", t)
	}()

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

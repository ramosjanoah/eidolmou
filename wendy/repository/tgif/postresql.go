package tgif

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/connection"
	"github.com/yanzay/tbot/v2"
	"net/http"
	"strconv"
	"time"
)

type GifPostgresql struct{}

func GetGifPostgresqlRepository() *GifPostgresql {
	return &GifPostgresql{}
}

func (g *GifPostgresql) Check() error {
	return nil
}

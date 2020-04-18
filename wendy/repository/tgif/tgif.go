package tgif

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/repository"
)

func GetGifRepository() repository.GifRepository {
	if config.SQLType == "MYSQL" {
		return GetGifSqlRepository()
	}

	panic("SqlDatabase value is nil")
}

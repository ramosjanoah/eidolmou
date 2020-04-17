package tgif

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
)

func GetGifRepository() ActionBotRepository {
	if config.SqlDatabase == "POSTGRESQL" {
		return GetGifPostgresqlRepository()
	}

	panic("SqlDatabase value is nil")
}

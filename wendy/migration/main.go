package main

import (
	"log"

	"fmt"
	"github.com/go-gormigrate/gormigrate"
	"github.com/jinzhu/gorm"
	"github.com/ramosjanoah/eidolmou/wendy/connection"
	"github.com/ramosjanoah/eidolmou/wendy/model/tgif"
)

func main() {
	// get connection first

	if connection.SQLConn == nil {
		panic("Sql Connection is nil when about to migrate")
	}

	m := gormigrate.New(connection.SQLConn, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "2019-04-18-1649",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&tgif.TGif{}).Error
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.DropTable("t_gifs").Error
			},
		},
	})

	err := m.Migrate()
	if err != nil {
		panic(fmt.Sprintf("Migration failed: %s", err.Error()))
	}

	log.Println("Migration success")
}

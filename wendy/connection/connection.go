package connection

import (
	"fmt"

	// for config
	"github.com/ramosjanoah/eidolmou/wendy/config"

	// for sql connection
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// for telegram bot
	"github.com/yanzay/tbot/v2"

	// for heartbeat
	"net/http"
	"time"
)

const telegramAPIURL = "https://api.telegram.org"

var (
	TbotClient *tbot.Client
	Heartbeat  *http.Client
	SQLConn    *gorm.DB
)

func init() {

	var err error

	if config.Platform == "TELEGRAM" {
		httpClient := http.Client{
			Timeout: 10 * time.Second,
		}
		TbotClient = tbot.NewClient(config.BotToken, &httpClient, telegramAPIURL)
	}

	if config.HeartbeatToggle {
		Heartbeat = &http.Client{Timeout: 5 * time.Second}
	}

	if config.SQLType == "MYSQL" {
		addInfo := "charset=utf8&parseTime=True&loc=Local"
		SQLConn, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%s)/%s?%s",
			config.SQLUser,
			config.SQLPassword,
			config.SQLHost,
			config.SQLPort,
			config.SQLDatabase,
			addInfo,
		))
		if err != nil {
			panic(fmt.Sprintf("Failed to establish MySQL connection, %s", err.Error()))
		}
		if err = SQLConn.DB().Ping(); err != nil {
			panic(fmt.Sprintf("Failed to ping MYSQL connection, %s", err.Error()))
		}
	}
}

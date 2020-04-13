package telegram

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/handler/http"
	"github.com/ramosjanoah/eidolmou/wendy/service"
	"github.com/yanzay/tbot/v2"
	"log"
	netHttp "net/http"
)

type TelegramBot struct {
	Token    string
	HttpPort string

	botClient *tbot.Client
	botServer *tbot.Server
}

func NewBot() *TelegramBot {
	config := config.GetConfig()
	t := &TelegramBot{
		Token:    config.BotToken,
		HttpPort: config.AppPort,
	}

	err := t.Initialize()
	if err != nil {
		panic(err)
	}

	return t
}

func (t *TelegramBot) ChatListen() {
	log.Println("Wendy is listening for your chat..")
	t.botServer.Start()
}

func (t *TelegramBot) HttpListen() {
	router := http.NewHttpRouter()
	log.Println(fmt.Sprintf("Wendy is listening to your HTTP request in :%s", t.HttpPort))
	log.Println(netHttp.ListenAndServe(":"+t.HttpPort, router))
}

func (t *TelegramBot) Initialize() error {
	t.botServer = tbot.New(t.Token)
	t.botClient = t.botServer.Client()

	// create handling message here
	t.initializeHandler()

	return nil
}

func (t *TelegramBot) initializeHandler() error {
	t.botServer.HandleMessage("/areyouok", decorate(t.areYouOK))

	return nil
}

func (t *TelegramBot) areYouOK(m *tbot.Message) error {
	result := service.AreYouOK()

	_, err := t.botClient.SendMessage(m.Chat.ID, *result.MessageResponse)
	if err != nil {
		return err
	}

	return nil
}

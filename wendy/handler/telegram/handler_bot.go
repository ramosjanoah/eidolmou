package telegram

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/handler/http"
	"github.com/ramosjanoah/eidolmou/wendy/service"
	"github.com/yanzay/tbot/v2"
	"log"
	netHttp "net/http"
	"time"
)

type HandlerBot struct {
	Token    string
	HttpPort string

	botServer *tbot.Server

	heartbeatClient *netHttp.Client
	heartbeatURL    string
}

func NewHandlerBot() *HandlerBot {
	t := &HandlerBot{
		Token:    config.BotToken,
		HttpPort: config.AppPort,
	}

	err := t.Initialize()
	if err != nil {
		panic(err)
	}

	return t
}

func (t *HandlerBot) ChatListen() {
	log.Println("Wendy is listening for your chat..")
	t.botServer.Start()
}

func (t *HandlerBot) HttpListen() {
	router := http.NewHttpRouter()

	log.Println(fmt.Sprintf("Wendy is listening to your HTTP request in :%s", t.HttpPort))
	log.Println(netHttp.ListenAndServe(":"+t.HttpPort, router))
}

func (t *HandlerBot) Heartbeat() {
	if !config.HeartbeatToggle {
		return
	}

	req, err := netHttp.NewRequest("GET", t.heartbeatURL, nil)
	if err != nil {
		panic("Heartbeat failed")
	}

	for {
		_, err = t.heartbeatClient.Do(req)
		if err != nil {
			panic(fmt.Sprintf("Heartbeat failed, %s", err.Error()))
		}
		time.Sleep(3 * time.Minute)
	}
}

func (t *HandlerBot) Initialize() error {
	t.botServer = tbot.New(t.Token)

	err := t.initializeBotHandler()
	if err != nil {
		return err
	}

	err = t.initializeHeartbeat()
	if err != nil {
		return err
	}

	return nil
}

func (t *HandlerBot) initializeBotHandler() error {
	// create handling message here

	t.botServer.HandleMessage(decorate("/areyouok", t.areYouOK))
	t.botServer.HandleMessage(decorate("/sendmegif", t.sendMeGif))

	return nil
}

func (t *HandlerBot) initializeHeartbeat() error {
	t.heartbeatURL = http.HeartbeatURL
	t.heartbeatClient = &netHttp.Client{}

	return nil
}

func (t *HandlerBot) areYouOK(m *tbot.Message) (err error) {
	_, err = service.AreYouOK(int64(m.From.ID))
	return err
}

func (t *HandlerBot) sendMeGif(m *tbot.Message) (err error) {
	return service.SendMeGif(int64(m.From.ID))
}

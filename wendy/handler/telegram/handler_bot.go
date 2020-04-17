package telegram

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	"github.com/ramosjanoah/eidolmou/wendy/handler/http"
	service "github.com/ramosjanoah/eidolmou/wendy/service"
	"github.com/yanzay/tbot/v2"
	"log"
	netHttp "net/http"
	"strconv"
	"strings"
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

	// see tgif_handler.go
	t.botServer.HandleMessage(decorate("/sendmegif", t.sendMeGif))

	// handle the message that
	t.botServer.HandleMessage(decorate("", func(m *tbot.Message) error {

		// handling document and get the command from caption
		if m.Document != nil {
			pattern, _ := parseCaption(m.Caption)
			switch pattern {
			case "/addGif":
				return t.addGif(m)
			}
		}

		return errors.DontUnderstandError()
	}))

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

func parseCaption(caption string) (string, map[string]string) {
	splittedCaption := strings.Split(caption, " ")

	command := ""
	args := map[string]string{}
	var key, value string

	for i, c := range splittedCaption {
		if i == 0 && c[0] == '/' {
			command = c
			continue
		}

		key, value = parseArg(c, strconv.Itoa(i))
		args[key] = value
	}

	return command, args
}

func parseArg(str string, defaultKey string) (key string, value string) {
	splittedStr := strings.Split(str, "=")
	if len(splittedStr) == 2 {
		return splittedStr[0], splittedStr[1]
	} else {
		return defaultKey, str
	}
}

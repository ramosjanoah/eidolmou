package telegram

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	service "github.com/ramosjanoah/eidolmou/wendy/service"
	"github.com/yanzay/tbot/v2"
	"log"
	"strconv"
	"strings"
)

type ChatHandler struct {
	Token     string
	botServer *tbot.Server
}

func NewChatHandler() *ChatHandler {
	t := &ChatHandler{
		Token: config.BotToken,
	}

	err := t.Initialize()
	if err != nil {
		panic(err)
	}

	return t
}

func (t *ChatHandler) ChatListen() {
	log.Println("Wendy is listening for your chat.")
	t.botServer.Start()
}

func (t *ChatHandler) Initialize() error {
	t.botServer = tbot.New(t.Token)

	err := t.initializeBotChatHandler()
	if err != nil {
		return err
	}

	return nil
}

func (t *ChatHandler) initializeBotChatHandler() error {
	// create handling message here

	t.botServer.HandleMessage(decorate("/areyouok", t.areYouOK))

	// see tgif_ChatHandler.go
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

func (t *ChatHandler) areYouOK(m *tbot.Message) (err error) {
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
	}

	return defaultKey, str
}

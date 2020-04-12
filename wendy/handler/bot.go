package handler

import (
	"github.com/ramosjanoah/eidolmou/wendy/config"
	tb "gopkg.in/tucnak/telebot.v2"
)

func NewBot() *tb.Bot {
	c := config.GetConfig()

	pref := tb.Settings{
		Token: c.BotToken,
		Poller: &tb.Webhook{
			Listen:   ":" + c.AppPort,
			Endpoint: &tb.WebhookEndpoint{PublicURL: c.URL},
		},
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		panic(err)
	}

	err = AddHandler(b)
	if err != nil {
		panic(err)
	}

	return b
}

func AddHandler(bot *tb.Bot) error {
	bot.Handle("/hello", func(m *tb.Message) {
		bot.Send(m.Sender, "Hi!")
	})

	return nil
}

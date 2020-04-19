package actionbot

import (
	"github.com/ramosjanoah/eidolmou/wendy/connection"
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	"github.com/ramosjanoah/eidolmou/wendy/model/message"
	"github.com/ramosjanoah/eidolmou/wendy/sctx"
	"github.com/yanzay/tbot/v2"
	"strconv"
)

type TelegramActionBot struct {
	BotClient *tbot.Client
}

func GetTelegramActionBot() *TelegramActionBot {
	return &TelegramActionBot{connection.TbotClient}
}

func (a *TelegramActionBot) SendString(sctx sctx.Context, targetID int64, str string) (err error) {
	select {
	case <-sctx.Done():
		return errors.TimeoutError()
	default:
	}
	ctxl := sctx.AddLog("TelegramActionBot", "SendString").
		WithInfo("parameter:targetID", targetID).
		WithInfo("parameter:str", str)
	defer func() {
		ctxl.EndWithError(err)
	}()

	_, err = a.BotClient.SendMessage(strconv.Itoa(int(targetID)), str)
	return err
}
func (a *TelegramActionBot) SendMessage(sctx sctx.Context, targetID int64, m *message.Message) (err error) {
	select {
	case <-sctx.Done():
		return errors.TimeoutError()
	default:
	}
	ctxl := sctx.AddLog("TelegramActionBot", "SendMessage").
		WithInfo("parameter:targetID", targetID).
		WithInfo("parameter:m.GetString()", m.GetString())
	defer func() {
		ctxl.EndWithError(err)
	}()

	_, err = a.BotClient.SendMessage(strconv.Itoa(int(targetID)), m.GetString(), tbot.OptParseModeMarkdown)
	return err
}

func (a *TelegramActionBot) SendAnimation(sctx sctx.Context, targetID int64, animationURL string) (err error) {
	select {
	case <-sctx.Done():
		return errors.TimeoutError()
	default:
	}
	ctxl := sctx.AddLog("TelegramActionBot", "SendAnimation").
		WithInfo("parameter:targetID", targetID).
		WithInfo("parameter:animationURL", animationURL)
	defer func() {
		ctxl.EndWithError(err)
	}()

	_, err = a.BotClient.SendAnimation(strconv.Itoa(int(targetID)), animationURL)
	return err
}

func (a *TelegramActionBot) SendAnimationFile(sctx sctx.Context, targetID int64, animationFilename string) (err error) {
	select {
	case <-sctx.Done():
		return errors.TimeoutError()
	default:
	}
	ctxl := sctx.AddLog("TelegramActionBot", "SendAnimationFile").
		WithInfo("parameter:targetID", targetID).
		WithInfo("parameter:animationFilename", animationFilename)
	defer func() {
		ctxl.EndWithError(err)
	}()

	_, err = a.BotClient.SendAnimationFile(strconv.Itoa(int(targetID)), animationFilename)
	return err
}

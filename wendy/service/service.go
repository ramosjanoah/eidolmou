package service

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"github.com/ramosjanoah/eidolmou/wendy/errors"
	"github.com/ramosjanoah/eidolmou/wendy/repository"
	"github.com/ramosjanoah/eidolmou/wendy/repository/actionbot"
	"github.com/ramosjanoah/eidolmou/wendy/sctx"
)

// list all the repository needed for this service
var ActionBot repository.ActionBotRepository

// initialize the repository
func init() {
	ActionBot = actionbot.GetActionBot()
}

func AreYouOK(sctx sctx.Context, targetID int64) (result string, err error) {
	select {
	case <-sctx.Done():
		return "", errors.TimeoutError()
	default:
	}
	ctxl := sctx.AddLog("service", "AreYouOK").WithInfo("parameter:targetID", targetID)
	defer func() {
		ctxl.EndWithError(err)
	}()

	// send message response for 'are you ok'
	err = ActionBot.SendString(sctx, targetID, AreYouOKResponseMsg)
	if err != nil {
		return "", err
	}

	return AreYouOKResponseMsg, nil
}

func PleaseCheckMe(sctx sctx.Context) (err error) {
	ctxl := sctx.AddLog("service", "PleaseCheckMe")
	defer func() {
		ctxl.EndWithError(err)
	}()

	if config.AdminID == 0 {
		return nil
	}

	// send check me message to admin
	return ActionBot.SendString(sctx, config.AdminID, PleaseCheckMeMsg)
}

func ErrorResponseCallback(sctx sctx.Context, targetID int64, err error) (callbackErr error) {
	ctxl := sctx.AddLog("service", "ErrorResponseCallback").WithInfo("parameter:targetID", targetID)
	defer func() {
		ctxl.EndWithError(callbackErr)
	}()

	return ActionBot.SendString(sctx, targetID, err.Error())
}

func sugarCoatError(err error) string {
	return fmt.Sprintf("Damn :(\nyou got some error: '%s'", err.Error())
}

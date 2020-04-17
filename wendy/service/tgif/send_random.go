package tgif

import (
	"fmt"
	"github.com/ramosjanoah/eidolmou/wendy/config"
	"math/rand"
	"time"
)

var WendyGifs = []string{
	"wendy-hi",
	"wendy-swag",
	"wendy-love-thumbs",
}

func SendMeGif(targetID int64) error {
	// pick random animation file
	rand.Seed(time.Now().Unix())
	pickedFile := fmt.Sprintf("%s/wendy/asset/%s.gif", config.CurrentDir, WendyGifs[rand.Int()%len(WendyGifs)])
	return ActionBot.SendAnimationFile(targetID, pickedFile)
}

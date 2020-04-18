package tgif

func SendMeGif(targetID int64) error {
	// get how many gifs are there

	randomTGif, err := TGifRepository.GetRandom()
	if err != nil {
		return err
	}

	return ActionBotRepository.SendAnimation(targetID, randomTGif.FileID)
}

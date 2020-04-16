package repository

type ActionBotRepository interface {
	SendMessage(targetID int64, message string) error
	SendAnimation(targetID int64, animationURL string) error
	SendAnimationFile(targetID int64, animiationFileName string) error
}

type GifRepository interface {
	Check() error
}

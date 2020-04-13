package service

type Service struct{}

var (
	AreYouOKResponse = "Hi! I'm wendy and I'm ok :)"
)

type Result struct {
	MessageResponse *string
}

func AreYouOK() Result {
	// There are no logic in AreYouOK so I'll

	return Result{
		MessageResponse: &AreYouOKResponse,
	}
}

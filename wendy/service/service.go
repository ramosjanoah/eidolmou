package service

type Service struct{}

var (
	AreYouOKResponse = "Hi! I'm ok, don't worry :)"
)

type Result struct {
	MessageResponse *string
}

func AreYouOK() Result {
	// There are no logic in AreYouOK so I'll just return the response

	return Result{
		MessageResponse: &AreYouOKResponse,
	}
}

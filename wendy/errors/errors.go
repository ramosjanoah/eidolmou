package errors

import (
	"fmt"
	"github.com/bukalapak/apinizer/response"
)

func DontUnderstandError() error {
	return response.CustomError{Message: "Sorry, I don't understand :(", HTTPCode: 403}
}

func MissingParameterError(field string) error {
	return response.CustomError{Message: fmt.Sprintf("Sorry, but you're missing this parameter: '%s'", field), HTTPCode: 400}
}

func PayloadMissingError() error {
	return response.CustomError{Message: "Sorry, but you sent the empty message", HTTPCode: 400}
}

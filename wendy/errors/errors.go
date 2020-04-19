package errors

import (
	"fmt"
)

type CustomError struct {
	Message            string
	SugarcoatedMessage string
	HTTPCode           int
}

func (c *CustomError) Error() string {
	return c.Message
}

func DontUnderstandError() error {
	return &CustomError{Message: "Sorry, I don't understand :(", HTTPCode: 403}
}

func MissingParameterError(field string) error {
	return &CustomError{Message: fmt.Sprintf("Sorry, but you're missing this parameter: '%s'", field), HTTPCode: 400}
}

func PayloadMissingError() error {
	return &CustomError{Message: "Sorry, but you sent the empty message", HTTPCode: 400}
}

func ErrorWhenInsertRow(instanceType string, detailedError error) error {
	return &CustomError{Message: fmt.Sprintf("Error when create %s: %s", instanceType, detailedError.Error()), HTTPCode: 400}
}

func NotFoundError(instanceType string) error {
	return &CustomError{Message: fmt.Sprintf("%s not found :(", instanceType), HTTPCode: 404}
}

func TimeoutError() error {
	return &CustomError{Message: "Timeout", HTTPCode: 408}
}

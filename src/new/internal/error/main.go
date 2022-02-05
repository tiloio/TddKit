package error

import "log"

type commonError struct {
	Error   error
	Message string
}

type Common = *commonError

func NewCommon(err error, message string) Common {
	return &commonError{Error: err, Message: message}
}

func (err Common) Fatal() {
	log.Fatal(err.Message, err.Error)
}

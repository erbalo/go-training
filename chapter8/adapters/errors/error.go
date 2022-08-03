package errors

import "fmt"

type error interface {
	Error() string
}

type InternalError struct {
	Message string
}

func (error *InternalError) Error() string {
	return fmt.Sprintf("error test %v", error.Message)
}

func NewInternal(message string) error {
	return &InternalError{
		Message: message,
	}
}

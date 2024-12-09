package utils

import "fmt"

type Error struct {
	message string
}

func (e Error) Error() error {
	return fmt.Errorf(e.message)
}

func (e Error) Errorf(args ...interface{}) error {
	return fmt.Errorf(e.message, args...)
}

var (
	ErrEmptyString = &Error{"empty hex string"}
	ErrUnknownChar = &Error{"invalid byte: %#U"}
	ErrOddLength   = &Error{"odd length hex string"}
)

var ErrTypesLength = &Error{"invalid bytes length got %v, expected %v"}

var (
	ErrUnsupportType = &Error{"type %v is not support"}
	ErrNaNFloat      = &Error{"NaN float value"}
)

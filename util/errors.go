package util

import "fmt"

var ErrRequired = fmt.Errorf("cannot be required")
var ErrInvalidEmail = fmt.Errorf("should be a valid email")
var ErrInvalidPhone = fmt.Errorf("should only contain numbers and should be of length 11")
var ErrInvalidAlphaNum = fmt.Errorf("should only contain letters and numbers")
var ErrInvalidAlpha = fmt.Errorf("should only contain letters")

type ErrInvalidLength error

func NewErrInvalidLength(length int) ErrInvalidLength {
	return fmt.Errorf("should be %v chars long", length)
}

type ErrMinLength error

func NewErrMinLength(length int) ErrMinLength {
	return fmt.Errorf("should be at least %v chars long", length)
}

type ErrMaxLength error

func NewErrMaxLength(length int) ErrMaxLength {
	return fmt.Errorf("should be less than %v chars long", length)
}

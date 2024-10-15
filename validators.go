package main

import (
	"FormValidator/util"
	"reflect"
	"regexp"
)

var alphaRe = regexp.MustCompile(`^[a-zA-Z]+$`)
var alphaNumRe = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
var phoneRe = regexp.MustCompile(`^\d{11}$`)
var emailRe = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Validator Generic data validator
type Validator interface {
	Validate(any) (bool, error)
}

type DefaultValidator struct{}

func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

type FormValidator struct {
	Min      int
	Max      int
	Length   int
	Alpha    bool
	AlphaNum bool
	Required bool
	Phone    bool
	Email    bool
}

func (v FormValidator) Validate(val any) (bool, error) {
	t := reflect.TypeOf(val)

	var l int
	if t.Name() == "int" {
		l = val.(int)
	} else {
		l = len(val.(string))
	}

	if v.Required && l == 0 {
		return false, util.ErrRequired
	}

	if v.Length > 0 && l != v.Length {
		return false, util.NewErrInvalidLength(v.Length)
	}

	if l < v.Min {
		return false, util.NewErrMinLength(v.Min)
	}

	if v.Max > 0 && l > v.Max {
		return false, util.NewErrMaxLength(v.Max)
	}

	if v.Alpha && !alphaRe.MatchString(val.(string)) {
		return false, util.ErrInvalidAlpha
	}

	if v.AlphaNum && !alphaNumRe.MatchString(val.(string)) {
		return false, util.ErrInvalidAlphaNum
	}

	if v.Phone && !phoneRe.MatchString(val.(string)) {
		return false, util.ErrInvalidPhone
	}

	if v.Email && !emailRe.MatchString(val.(string)) {
		return false, util.ErrInvalidEmail
	}

	return true, nil
}

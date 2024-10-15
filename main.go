package main

import (
	"fmt"
	"reflect"
	"strings"
)

const tagName = "form"

type UserDetails struct {
	FirstName   string `form:"required,min=5,alpha"`
	LastName    string `form:"required,min=5,alpha"`
	Email       string `form:"required,email"`
	PhoneNumber string `form:"required,phone"`
	Age         int    `form:"required,min=12"`
}

func getValidatorFromTag(tag string) Validator {
	validator := FormValidator{}
	for _, option := range strings.Split(tag, ",") {
		if strings.HasPrefix(option, "min") {
			_, _ = fmt.Sscanf(option, "min=%d", &validator.Min)
		} else if strings.HasPrefix(option, "max") {
			_, _ = fmt.Sscanf(option, "max=%d", &validator.Max)
		} else if strings.HasPrefix(option, "len") {
			_, _ = fmt.Sscanf(option, "len=%d", &validator.Length)
		} else {
			validator.Required = option == "required"
			validator.Phone = option == "phone"
			validator.Alpha = option == "alpha"
			validator.AlphaNum = option == "alphanum"
			validator.Email = option == "email"
		}
	}
	return validator
}

func validateStruct(s any) error {
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagName)

		if tag == "" {
			continue
		}

		validator := getValidatorFromTag(tag)

		_, err := validator.Validate(v.Field(i).Interface())

		if err != nil {
			return fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error())
		}

	}
	return nil
}

func main() {

}

package main

import (
	"FormValidator/util"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestStringValidator(t *testing.T) {
	testCases := []struct {
		name       string
		buildStubs func(t *testing.T)
	}{
		{
			name: "Default",
			buildStubs: func(t *testing.T) {
				validator := FormValidator{}

				valid, err := validator.Validate("")
				require.NoError(t, err)
				require.True(t, valid)
			},
		},
		{
			name: "Required",
			buildStubs: func(t *testing.T) {
				validator := FormValidator{
					Required: true,
				}

				valid, err := validator.Validate(util.RandomString(5))
				require.NoError(t, err)
				require.True(t, valid)

				valid, err = validator.Validate("")
				require.Error(t, err)
				require.EqualError(t, err, util.ErrRequired.Error())
				require.False(t, valid)
			},
		},
		{
			name: "Length/Ok",
			buildStubs: func(t *testing.T) {
				length := 7
				validator := FormValidator{
					Length: length,
				}

				valid, err := validator.Validate(util.RandomString(length))
				require.NoError(t, err)
				require.True(t, valid)

				valid, err = validator.Validate(util.RandomString(length - 1))
				require.Error(t, err)
				require.EqualError(t, err, util.NewErrInvalidLength(length).Error())
				require.False(t, valid)

				valid, err = validator.Validate(util.RandomString(length + 1))
				require.Error(t, err, util.NewErrInvalidLength(length).Error())
				require.False(t, valid)
			},
		},
		{
			name: "Min",
			buildStubs: func(t *testing.T) {
				minLen := 3
				validator := FormValidator{
					Min: minLen,
				}

				valid, err := validator.Validate(util.RandomString(minLen))
				require.NoError(t, err)
				require.True(t, valid)

				valid, err = validator.Validate(util.RandomString(minLen - 1))
				require.Error(t, err)
				require.EqualError(t, err, util.NewErrMinLength(minLen).Error())
				require.False(t, valid)

				valid, err = validator.Validate(minLen)
				require.NoError(t, err)
				require.True(t, valid)
			},
		},
		{
			name: "Max",
			buildStubs: func(t *testing.T) {
				maxLen := 10
				validator := FormValidator{
					Max: maxLen,
				}

				valid, err := validator.Validate(util.RandomString(maxLen))
				require.NoError(t, err)
				require.True(t, valid)

				valid, err = validator.Validate(util.RandomString(maxLen + 1))
				require.Error(t, err)
				require.EqualError(t, err, util.NewErrMaxLength(maxLen).Error())
				require.False(t, valid)
			},
		},
		{
			name: "Alpha",
			buildStubs: func(t *testing.T) {
				length := 5
				validator := FormValidator{
					Alpha: true,
				}

				valid, err := validator.Validate(util.RandomString(length))
				require.NoError(t, err)
				require.True(t, valid)

				valid, err = validator.Validate(util.RandomAlphaNumString(length))
				require.Error(t, err)
				require.EqualError(t, err, util.ErrInvalidAlpha.Error())
				require.False(t, valid)
			},
		},
		{
			name: "AlphaNum",
			buildStubs: func(t *testing.T) {
				length := 5
				validator := FormValidator{
					AlphaNum: true,
				}

				valid, err := validator.Validate(util.RandomString(length))
				require.NoError(t, err)
				require.True(t, valid)

				valid, err = validator.Validate(util.RandomAlphaNumString(length))
				require.NoError(t, err)
				require.True(t, valid)

				valid, err = validator.Validate(util.RandomAlphaNumString(length) + "@")
				require.Error(t, err)
				require.EqualError(t, err, util.ErrInvalidAlphaNum.Error())
				require.False(t, valid)
			},
		},
		{
			name: "Phone",
			buildStubs: func(t *testing.T) {
				length := 11
				validator := FormValidator{
					Phone: true,
				}

				valid, err := validator.Validate(util.RandomNumString(length))
				require.NoError(t, err)
				require.True(t, valid)

				valid, err = validator.Validate(util.RandomString(length))
				require.Error(t, err)
				require.EqualError(t, err, util.ErrInvalidPhone.Error())
				require.False(t, valid)

				valid, err = validator.Validate(util.RandomAlphaNumString(length))
				require.Error(t, err)
				require.EqualError(t, err, util.ErrInvalidPhone.Error())
				require.False(t, valid)

				valid, err = validator.Validate(util.RandomNumString(length - 1))
				require.Error(t, err)
				require.EqualError(t, err, util.ErrInvalidPhone.Error())
				require.False(t, valid)
			},
		},
		{
			name: "Email",
			buildStubs: func(t *testing.T) {
				validator := FormValidator{
					Email: true,
				}
				fmt.Println(util.RandomEmail())
				valid, err := validator.Validate(util.RandomEmail())
				require.NoError(t, err)
				require.True(t, valid)

				valid, err = validator.Validate(util.RandomString(9))
				require.Error(t, err)
				require.EqualError(t, err, util.ErrInvalidEmail.Error())
				require.False(t, valid)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.buildStubs)
	}
}

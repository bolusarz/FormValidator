package main

import (
	"FormValidator/util"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestValidation(t *testing.T) {

	testCases := []struct {
		name    string
		runTest func(t *testing.T)
	}{
		{
			name: "EmptyInit",
			runTest: func(t *testing.T) {
				user := UserDetails{}
				err := validateStruct(user)

				require.Error(t, err)
			},
		},
		{
			name: "Ok",
			runTest: func(t *testing.T) {
				user := UserDetails{
					FirstName:   util.RandomString(6),
					LastName:    util.RandomString(6),
					Email:       util.RandomEmail(),
					PhoneNumber: util.RandomNumString(11),
					Age:         int(util.RandomInt(12, 90)),
				}
				err := validateStruct(user)

				require.NoError(t, err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, tc.runTest)
	}
}

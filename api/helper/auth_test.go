package helper_test

import (
	"testing"

	"github.com/AdhityaRamadhanus/gopatrol/api/helper"
	"github.com/stretchr/testify/assert"
)

func TestParseAuth(t *testing.T) {
	testCases := []struct {
		Scheme        string
		Header        string
		ExpectedCreds string
		ExpectError   bool
	}{
		{
			Scheme:        "Bearer",
			Header:        "Bearer creds",
			ExpectedCreds: "creds",
			ExpectError:   false,
		},
		{
			Scheme:        "Bearer",
			Header:        "Basic creds",
			ExpectedCreds: "",
			ExpectError:   true,
		},
		{
			Scheme:        "Basic",
			Header:        "Basic creds",
			ExpectedCreds: "creds",
			ExpectError:   false,
		},
	}

	for _, test := range testCases {
		cred, err := helper.ParseAuthorizationHeader(test.Header, test.Scheme)
		if test.ExpectError {
			assert.Error(t, err, "Error parsing parse authorization header")
			assert.Empty(t, cred, "Parsed value is wrong")
		} else {
			assert.Empty(t, err, "Error parsing parse authorization header")
			assert.Equal(t, test.ExpectedCreds, cred, "Parsed value is wrong")
		}
	}
}

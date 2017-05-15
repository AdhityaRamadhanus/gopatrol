package helper

import (
	"strings"

	"github.com/pkg/errors"
)

func ParseAuthorizationHeader(authHeader, scheme string) (cred string, err error) {
	splittedHeader := strings.Split(authHeader, " ")
	if len(splittedHeader) != 2 {
		return "", errors.New("Cannot parse authorization header")
	}
	parsedScheme := splittedHeader[0]
	if scheme != parsedScheme {
		return "", errors.New("Unexpected Scheme, expected " + scheme)
	}
	return splittedHeader[1], nil
}

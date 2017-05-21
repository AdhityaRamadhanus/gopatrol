package helper

import (
	"strings"
)

func Slugify(text string) string {
	splittedStrings := []string{}
	for _, v := range strings.Split(text, " ") {
		splittedStrings = append(splittedStrings, strings.ToLower(v))
	}
	return strings.Join(splittedStrings, "-")
}

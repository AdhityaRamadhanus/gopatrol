package helper_test

import (
	"log"
	"testing"

	"github.com/AdhityaRamadhanus/gopatrol/api/helper"
	"github.com/stretchr/testify/assert"
)

func TestSlugify(t *testing.T) {
	testCases := []struct {
		Input          string
		ExpectedOutput string
	}{
		{
			Input:          "Local Redis",
			ExpectedOutput: "local-redis",
		},
		{
			Input:          "Redis",
			ExpectedOutput: "redis",
		},
	}

	for _, test := range testCases {
		log.Println("Slugigy", test.Input)
		output := helper.Slugify(test.Input)
		assert.Equal(t, test.ExpectedOutput, output, "Parsed value is wrong")
	}
}

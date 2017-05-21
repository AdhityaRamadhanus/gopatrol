package helper_test

import (
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
		output := helper.Slugify(test.Input)
		assert.Equal(t, test.ExpectedOutput, output, "Parsed value is wrong")
	}
}

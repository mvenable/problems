package strings_test

import (
	"fmt"
	"testing"

	"github.com/mvenable/problems/go/internal/strings"
	"github.com/stretchr/testify/assert"
)

var (
	matrix = []struct {
		Name     string
		Input    string
		Expected string
	}{
		{
			Name:     "negative",
			Input:    "54321",
			Expected: "12345",
		},
	}
)

func TestReverse(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := strings.Reverse(test.Input)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

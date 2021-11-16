package breaking_chocolate_test

import (
	"asdf/go/codewars/breaking_chocolate"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	Input    Params
	Expected int
}

type Params struct {
	Length int
	Width  int
}

var matrix = []TestCase{
	{
		Name: "Example Test 1",
		Input: Params{
			Length: 5,
			Width:  5,
		},
		Expected: 24,
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := breaking_chocolate.BreakingChocolate(test.Input.Length, test.Input.Width)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

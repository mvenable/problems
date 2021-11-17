package count_digit_test

import (
	"fmt"
	"testing"

	"github.com/mvenable/problems/go/codewars/count_digit"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string
	Input    Params
	Expected int
}

type Params struct {
	Integer int
	Digit   int
}

var matrix = []TestCase{
	{
		Name: "Example Test 1",
		Input: Params{
			Integer: 10,
			Digit:   1,
		},
		Expected: 4,
	},
	{
		Name: "Example Test 2",
		Input: Params{
			Integer: 25,
			Digit:   1,
		},
		Expected: 11,
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := count_digit.CountDigit(test.Input.Integer, test.Input.Digit)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

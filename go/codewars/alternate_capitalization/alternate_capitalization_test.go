package alternate_capitalization_test

import (
	"fmt"
	"testing"

	"github.com/mvenable/problems/go/codewars/alternate_capitalization"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string
	Input    string
	Expected Result
}

type Result struct {
	EvenCapitalized string
	OddCapitalized  string
}

var matrix = []TestCase{
	{
		Name:  "Example Test 1",
		Input: "abcdef",
		Expected: Result{
			EvenCapitalized: "AbCdEf",
			OddCapitalized:  "aBcDeF",
		},
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		even, odd := alternate_capitalization.Capitalize(test.Input)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, Result{EvenCapitalized: even, OddCapitalized: odd}, msg)
	}
}

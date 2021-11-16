package abbreviate_name_test

import (
	"asdf/go/codewars/abbreviate_name"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	Input    string
	Expected string
}

var matrix = []TestCase{
	{
		Name:     "Example Test 1",
		Input:    "Same Harris",
		Expected: "S.H",
	},
	{
		Name:     "Example Test 2",
		Input:    "Patrick feeney",
		Expected: "P.F",
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := abbreviate_name.AbbreviateName(test.Input)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

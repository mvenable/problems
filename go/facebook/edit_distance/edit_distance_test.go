package edit_distance_test

import (
	"asdf/go/facebook/edit_distance"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	Input    Params
	Expected bool
}

type Params struct {
	String1 string
	String2 string
}

var matrix = []TestCase{
	{
		Name:     "Completely Different",
		Input: Params{
			String1:  "cat",
			String2:  "dog",
		},
		Expected: false,
	},
	{
		Name:     "Second Suffix",
		Input: Params{
			String1:  "cat",
			String2:  "cats",
		},
		Expected: true,
	},
	{
		Name:     "Middle Different",
		Input: Params{
			String1:  "cat",
			String2:  "cut",
		},
		Expected: true,
	},
	{
		Name:     "Middle Different - Length",
		Input: Params{
			String1:  "cat",
			String2:  "cast",
		},
		Expected: true,
	},
	{
		Name:     "Missing Prefix",
		Input: Params{
			String1:  "cat",
			String2:  "at",
		},
		Expected: true,
	}, {
		Name:     "Swapped Letters",
		Input: Params{
			String1:  "cat",
			String2:  "act",
		},
		Expected: false,
	},
	{
		Name:     "UTF-8 rune length",
		Input: Params{
			String1:  "日at",
			String2:  "at",
		},
		Expected: true,
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := edit_distance.OneEditApart(test.Input.String1, test.Input.String2)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

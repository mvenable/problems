package consecutive_strings_test

import (
	"asdf/go/codewars/consecutive_strings"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	Input    Params
	Expected string
}

type Params struct {
	StringArray []string
	Consecutive int
}

var matrix = []TestCase{
	{
		Name: "Example Test 1",
		Input: Params{
			StringArray: []string{"tree", "foling", "trashy", "blue", "abcdef", "uvwxyz"},
			Consecutive: 2,
		},
		Expected: "folingtrashy",
	},
	{
		Name: "Example Test 2",
		Input: Params{
			StringArray: []string{"zone", "abigail", "theta", "form", "libe", "zas", "theta", "abigail"},
			Consecutive: 2,
		},
		Expected: "abigailtheta",
	},
	{
		Name: "Example Test 3",
		Input: Params{
			StringArray: []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"},
			Consecutive: 1,
		},
		Expected: "oocccffuucccjjjkkkjyyyeehh",
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := consecutive_strings.LongestConsecutive(test.Input.StringArray, test.Input.Consecutive)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

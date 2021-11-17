package look_and_say_test

import (
	"fmt"
	"testing"

	"github.com/mvenable/problems/go/facebook/look_and_say"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string
	Input    int
	Expected []string
}

var matrix = []TestCase{
	{
		Name:  "The One Test - Take it to Mordor!",
		Input: 10,
		Expected: []string{
			"1",
			"11",
			"21",
			"1211",
			"111221",
			"312211",
			"13112221",
			"1113213211",
			"31131211131221",
			"13211311123113112211",
		},
	},
	{
		Name:  "Validate Memoization",
		Input: 4,
		Expected: []string{
			"1",
			"11",
			"21",
			"1211",
		},
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := look_and_say.LookAndSay(test.Input)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

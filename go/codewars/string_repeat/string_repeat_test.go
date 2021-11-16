package string_repeat_test

import (
	"asdf/go/codewars/string_repeat"
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
	Repetitions int
	String      string
}

var matrix = []TestCase{
	{
		Name:     "Example Test",
		Input:    Params{Repetitions: 3, String: "test"},
		Expected: "testtesttest",
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := string_repeat.RepeatStr(test.Input.Repetitions, test.Input.String)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

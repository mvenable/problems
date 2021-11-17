package in_array_test

import (
	"fmt"
	"testing"

	"github.com/mvenable/problems/go/codewars/in_array"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string
	Input    Params
	Expected []string
}

type Params struct {
	Set1 []string
	Set2 []string
}

var matrix = []TestCase{
	{
		Name: "Example Test 1",
		Input: Params{
			Set1: []string{"live", "arp", "strong"},
			Set2: []string{"lively", "alive", "harp", "sharp", "armstrong"},
		},
		Expected: []string{"arp", "live", "strong"},
	},
	{
		Name: "Example Test 2",
		Input: Params{
			Set1: []string{"cod", "code", "wars", "ewar", "ar"},
			Set2: []string{},
		},
		Expected: []string{},
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := in_array.InArray(test.Input.Set1, test.Input.Set2)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

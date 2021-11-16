package phone_number_test

import (
	"asdf/go/codewars/phone_number"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	Input    [10]uint
	Expected string
}

var matrix = []TestCase{
	{
		Name:     "Example Test 1",
		Input:    [10]uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		Expected: "(012) 345-6789",
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := phone_number.CreatePhoneNumber(test.Input)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

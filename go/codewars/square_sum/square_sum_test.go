package square_sum_test

import (
	"fmt"
	"testing"

	"github.com/mvenable/problems/go/codewars/square_sum"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string
	Input    []int
	Expected int
}

var matrix = []TestCase{
	{
		Name:     "Example Test",
		Input:    []int{1, 2, 2},
		Expected: 9,
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := square_sum.SquareSum(test.Input)
		msg := fmt.Sprintf("%s - %v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

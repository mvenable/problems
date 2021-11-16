package is_triangle_test

import (
	"asdf/go/codewars/is_triangle"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	Input    [3]int
	Expected bool
}

var matrix = []TestCase{
	{
		Name: "Example Test 1",
		Input: [3]int{5, 1, 2},
		Expected: false,
	},
	{
		Name: "Example Test 2",
		Input: [3]int{2, 5, 1},
		Expected: false,
	},
	{
		Name: "Example Test 3",
		Input: [3]int{4, 2, 3},
		Expected: true,
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := is_triangle.IsTriangle(test.Input[0], test.Input[1], test.Input[2])
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

package my_strings_test

import (
	"asdf/go/my_strings"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	Input    Params
	Expected []int
}

type Params struct {
	Array []int
	Shift int
}

var (
	matrixRotateRight = []TestCase{
		{
			Name: "negative",
			Input: Params{
				Array: []int{1, 2, 3, 4, 5},
				Shift: -1,
			},
			Expected: []int{2, 3, 4, 5, 1},
		},
		{
			Name: "zero",
			Input: Params{
				Array: []int{1, 2, 3, 4, 5},
				Shift: 0,
			},
			Expected: []int{1, 2, 3, 4, 5},
		},
		{
			Name: "right one",
			Input: Params{
				Array: []int{1, 2, 3, 4, 5},
				Shift: 1,
			},
			Expected: []int{5, 1, 2, 3, 4},
		},
		{
			Name: "right four",
			Input: Params{
				Array: []int{1, 2, 3, 4, 5},
				Shift: 4,
			},
			Expected: []int{2, 3, 4, 5, 1},
		},
		{
			Name: "wrap around",
			Input: Params{
				Array: []int{1, 2, 3, 4, 5},
				Shift: 8,
			},
			Expected: []int{3, 4, 5, 1, 2},
		},
	}
	matrixRotateLeft = []TestCase{
		{
		Name: "negative",
			Input: Params{
				Array: []int{1, 2, 3, 4, 5},
				Shift: -1,
			},
			Expected: []int{5, 1, 2, 3, 4},
		},
		{
		Name: "zero",
			Input: Params{
				Array: []int{1, 2, 3, 4, 5},
				Shift: 0,
			},
			Expected: []int{1, 2, 3, 4, 5},
		},
		{
		Name: "left one",
			Input: Params{
				Array: []int{1, 2, 3, 4, 5},
				Shift: 1,
			},
			Expected: []int{2, 3, 4, 5, 1},
		},
		{
		Name: "left four",
			Input: Params{
				Array: []int{1, 2, 3, 4, 5},
				Shift: 4,
			},
			Expected: []int{5, 1, 2, 3, 4},
		},
		{
		Name: "wrap around",
			Input: Params{
				Array: []int{1, 2, 3, 4, 5},
				Shift: 8,
			},
			Expected: []int{4, 5, 1, 2, 3},
		},
	}
)


func Test(t *testing.T) {
	for i, test := range matrixRotateRight {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := my_strings.RotateRight(test.Input.Array, test.Input.Shift)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
	for i, test := range matrixRotateLeft {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := my_strings.RotateLeft(test.Input.Array, test.Input.Shift)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}
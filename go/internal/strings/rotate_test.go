package strings_test

import (
	"fmt"
	"testing"

	"github.com/mvenable/problems/go/internal/strings"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string
	Input    Params
	Expected string
}

type Params struct {
	Array string
	Shift int
}

var (
	matrixRotateRight = []TestCase{
		{
			Name: "negative",
			Input: Params{
				Array: "12345",
				Shift: -1,
			},
			Expected: "23451",
		},
		{
			Name: "zero",
			Input: Params{
				Array: "12345",
				Shift: 0,
			},
			Expected: "12345",
		},
		{
			Name: "right one",
			Input: Params{
				Array: "12345",
				Shift: 1,
			},
			Expected: "51234",
		},
		{
			Name: "right four",
			Input: Params{
				Array: "12345",
				Shift: 4,
			},
			Expected: "23451",
		},
		{
			Name: "wrap around",
			Input: Params{
				Array: "12345",
				Shift: 8,
			},
			Expected: "34512",
		},
	}
	matrixRotateLeft = []TestCase{
		{
			Name: "negative",
			Input: Params{
				Array: "12345",
				Shift: -1,
			},
			Expected: "51234",
		},
		{
			Name: "zero",
			Input: Params{
				Array: "12345",
				Shift: 0,
			},
			Expected: "12345",
		},
		{
			Name: "left one",
			Input: Params{
				Array: "12345",
				Shift: 1,
			},
			Expected: "23451",
		},
		{
			Name: "left four",
			Input: Params{
				Array: "12345",
				Shift: 4,
			},
			Expected: "51234",
		},
		{
			Name: "wrap around",
			Input: Params{
				Array: "12345",
				Shift: 8,
			},
			Expected: "45123",
		},
	}
)

func TestRotate(t *testing.T) {
	for i, test := range matrixRotateRight {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := strings.RotateRight(test.Input.Array, test.Input.Shift)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
	for i, test := range matrixRotateLeft {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := strings.RotateLeft(test.Input.Array, test.Input.Shift)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

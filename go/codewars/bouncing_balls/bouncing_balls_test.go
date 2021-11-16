package bouncing_balls_test

import (
	"asdf/go/codewars/bouncing_balls"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	Name     string
	Input    Params
	Expected int
}

type Params struct {
	Height float64
	Bounce float64
	Window float64
}

var matrix = []TestCase{
	{
		Name: "Example Test 1",
		Input: Params{
			Height: 3,
			Bounce: 0.66,
			Window: 1.5,
		},
		Expected: 3,
	},
	{
		Name: "Example Test 2",
		Input: Params{
			Height: 40,
			Bounce: 0.4,
			Window: 10,
		},
		Expected: 3,
	},
	{
		Name: "Example Test 3",
		Input: Params{
			Height: 10,
			Bounce: 0.6,
			Window: 10,
		},
		Expected: -1,
	},
	{
		Name: "Example Test 4",
		Input: Params{
			Height: 40,
			Bounce: 1,
			Window: 10,
		},
		Expected: -1,
	},
	{
		Name: "Example Test 5",
		Input: Params{
			Height: 5,
			Bounce: -1,
			Window: 1.5,
		},
		Expected: -1,
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		actual := bouncing_balls.BouncingBall(test.Input.Height, test.Input.Bounce, test.Input.Window)
		msg := fmt.Sprintf("%s - %+v", name, test.Input)
		assert.Equal(t, test.Expected, actual, msg)
	}
}

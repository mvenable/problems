package reverse_or_rotate_test

import (
	"fmt"
	"testing"

	"github.com/mvenable/problems/go/codewars/reverse_or_rotate"
	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name     string
	Input    Params
	Expected Response
}

type Params struct {
	String    string
	ChunkSize int
}

type Response struct {
	String string
	Error  error
}

var matrix = []TestCase{
	{
		Name: "Example Test",
		Input: Params{
			String:    "123456987654",
			ChunkSize: 6,
		},
		Expected: Response{
			String: "234561876549",
		},
	},
	{
		Name: "Example Test 2",
		Input: Params{
			String:    "123456987653",
			ChunkSize: 6,
		},
		Expected: Response{
			String: "234561356789",
		},
	},
}

func Test(t *testing.T) {
	for i, test := range matrix {
		name := test.Name
		if name == "" {
			name = fmt.Sprintf("Test %d", i)
		}

		str, err := reverse_or_rotate.ReverseOrRotate(test.Input.String, test.Input.ChunkSize)
		msg := fmt.Sprintf("%s - %v", name, test.Input)
		assert.Equal(t, test.Expected, Response{
			String: str,
			Error:  err,
		}, msg)
	}
}

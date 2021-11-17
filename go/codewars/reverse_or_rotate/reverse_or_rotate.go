package reverse_or_rotate

import (
	"math"
	"strconv"

	"github.com/mvenable/problems/go/internal/strings"
	"github.com/pkg/errors"
)

/*

ReverseOrRotate

The input is a string str of digits. Cut the string into chunks (a chunk here is a substring of the initial string)
of size sz (ignore the last chunk if its size is less than sz).

If a chunk represents an integer such as the sum of the cubes of its digits is divisible by 2, reverse that chunk;
otherwise rotate it to the left by one position. Put together these modified chunks and return the result as a string.

*/
func ReverseOrRotate(str string, chunkSize int) (string, error) {
	// While the directions specifically call out checking str being empty that case
	// is already covered by the two checks present.
	if chunkSize <= 0 || chunkSize > len(str) {
		return "", nil
	}

	out := ""
	for i := 0; i < len(str)/chunkSize; i++ {
		chunk := str[i*chunkSize : i*chunkSize+chunkSize]
		chunkCubeSum, err := cubeSum(chunk)
		if err != nil {
			return "", errors.Wrap(err, "failed to determine chunk cube sum")
		}
		if chunkCubeSum%2 == 0 {
			out += strings.Reverse(chunk)
		} else {
			out += strings.RotateLeft(chunk, 1)
		}
	}

	return out, nil
}

func cubeSum(str string) (int, error) {
	total := 0

	for _, char := range str {
		val, err := strconv.Atoi(string(char))
		if err != nil {
			return 0, errors.New("string was not numeric")
		}
		total += int(math.Pow(float64(val), 3))
	}

	return total, nil
}

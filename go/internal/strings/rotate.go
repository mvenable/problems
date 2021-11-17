package strings

import "math/big"

/*
RotateRight

Assuming an array is circular, the end links to the beginning, rotate an array right by `shift` positions.
*/
func RotateRight(str string, shift int) string {
	array := []rune(str)
	var out = make([]rune, len(array))
	for i := 0; i < len(array); i++ {
		out[i] = array[mod(i+len(array)-shift, len(array))]
	}
	return string(out)
}

/*
RotateLeft

Assuming an array is circular, the end links to the beginning, rotate an array left by `shift` positions.
*/
func RotateLeft(str string, shift int) string {
	array := []rune(str)
	// using make since the length is known.  saves potential memory reallocations.
	var out = make([]rune, len(array))
	for i := 0; i < len(array); i++ {
		out[i] = array[mod(i+shift, len(array))]
	}
	return string(out)
}

/*
  Today I learned that the functionality of the modulo operator
  is not equivalent between languages.  Explaination is available
  at: https://github.com/golang/go/issues/448#issuecomment-66049769

  Rather than implementing an algorithm myself, I've imported math.Big
  which has implemented the euclidean modulo for us.  Feels very messy.
*/
func mod(a, b int) int64 {
	c := big.NewInt(0)
	c.Mod(big.NewInt(int64(a)), big.NewInt(int64(b)))
	return c.Int64()
}

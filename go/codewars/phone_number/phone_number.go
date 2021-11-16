package phone_number

import (
	"fmt"
	"strconv"
)

/*
CreatePhoneNumber

Write a function that accepts an array of 10 integers (between 0 and 9), that returns a string of those numbers in the
form of a phone number.
*/
func CreatePhoneNumber(numbers [10]uint) string {
	// gather numbers into logical groups to make the string format more readable. It would be easy
	// enough to just print all the numbers individually.
	areaCode := uintSliceToStr(numbers[:3])
	firstThree := uintSliceToStr(numbers[3:6])
	lastFour := uintSliceToStr(numbers[6:])

	return fmt.Sprintf("(%s) %s-%s", areaCode, firstThree, lastFour)
}

// A different approach that could be taken here is to add the uint values together and use %d or %f.
// aka  []uint{1, 2, 3} == 100 + 20 + 3 = 123.  I used the string concatenation approach for readability,
// in liu of potential performance improvements.
func uintSliceToStr(in []uint) string {
	out := ""
	for _, num := range in {
		out += strconv.FormatUint(uint64(num), 10)
	}
	return out
}

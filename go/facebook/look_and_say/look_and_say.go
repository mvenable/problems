package look_and_say

import (
	"strconv"
)

/*
Implement a function that returns the Look and Say (Morris Number Sequence) of length X.

To generate a member of the sequence from the previous member, read off the digits of the previous member, counting
the number of digits in groups of the same digit. For example:

1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
1211 is read off as "one 1, one 2, then two 1s" or 111221.
111221 is read off as "three 1s, two 2s, then one 1" or 312211.

*/

// optimizing for performance of multiple calls.
// will slightly increase memory usage and could
// present itself as a very slow "memory leak"
// if consistently called with larger counts.
//
// used a map here because I am lazy and did not
// want to constantly length check before access.
var linesMemoized = map[int]string{}

func LookAndSay(iterations int) []string {
	var lastLine string
	var lines []string

	for idx := 0; idx < iterations; idx++ {
		if linesMemoized[idx] != "" {
			lastLine = linesMemoized[idx]
		} else {
			lastLine = nextLine(lastLine)
			linesMemoized[idx] = lastLine
		}
		lines = append(lines, lastLine)
	}

	return lines
}

func nextLine(currentLine string) string {
	if currentLine == "" {
		return "1"
	}

	out := ""
	numbers := strToIntArray(currentLine)
	count := 1
	lastNumber := numbers[0]

	for _, number := range numbers[1:] {
		if number == lastNumber {
			count++
		} else {
			out = out + strconv.Itoa(count) + strconv.Itoa(lastNumber)
			lastNumber = number
			count = 1
		}
	}

	return out + strconv.Itoa(count) + strconv.Itoa(lastNumber)
}

func strToIntArray(str string) []int {
	var intArray []int
	for _, val := range str {
		intArray = append(intArray, int(val-'0'))
	}
	return intArray
}

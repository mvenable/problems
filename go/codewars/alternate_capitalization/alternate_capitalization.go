package alternate_capitalization

import "unicode"

/*
Capitalize

Given a string, capitalize the letters that occupy even indexes and odd indexes separately, and return as shown below.
Index 0 will be considered even.

The input will be a lowercase string with no spaces.

NOTE:  The function signature provided returned []string.  This ambiguity in expected response size bothers me.  I
       modified it to specifically return two strings.
*/
func Capitalize(originalString string) (string, string) {
	evenStr := []rune(originalString)
	oddStr := []rune(originalString)

	i := 0
	for pos, char := range originalString {
		upperCaseRune := unicode.ToUpper(char)
		if i%2 == 0 {
			evenStr[pos] = upperCaseRune
		} else {
			oddStr[pos] = upperCaseRune
		}
		i++
	}
	return string(evenStr), string(oddStr)
}

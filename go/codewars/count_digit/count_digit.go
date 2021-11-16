package count_digit

import (
	"strconv"
	"strings"
)

/*
CountDigit

Take an integer n (n >= 0) and a digit d (0 <= d <= 9) as an integer.

Square all numbers k (0 <= k <= n) between 0 and n.

Count the numbers of digits d used in the writing of all the k**2.

Call nb_dig (or nbDig or ...) the function taking n and d as parameters and returning this count.
*/
func CountDigit(maxSquareRoot int, digit int) int {
	if maxSquareRoot < 0 || digit < 0 || digit > 9 {
		return -1
	}

	count := 0
	strDigit := strconv.Itoa(digit)
	for i := 0; i <= maxSquareRoot; i++ {
		count += strings.Count(strconv.Itoa(i*i), strDigit)
	}

	return count
}

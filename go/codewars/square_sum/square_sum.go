package square_sum

/*
SquareSum

Squares each number passed into it and then sums the results together.
For example, for [1, 2, 2] it should return 9 because 1^2 + 2^2 + 2^2 = 9.
*/
func SquareSum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number * number
	}
	return sum
}

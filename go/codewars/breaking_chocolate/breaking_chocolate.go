package breaking_chocolate

/*
BreakingChocolate

Your task is to split the chocolate bar of given dimension n x m into small squares. Each square is of size 1x1 and
unbreakable. Implement a function that will return minimum number of breaks needed.

For example if you are given a chocolate bar of size 2 x 1 you can split it to single squares in just one break, but
for size 3 x 1 you must do two breaks.

If input data is invalid you should return 0 (as in no breaks are needed if we do not have any chocolate to split).
Input will always be a non-negative integer.

*/
func BreakingChocolate(length, width int) int {
	// guard statement belongs at the top.  yes, we could remove one check by only checking width
	// after variable re-assignment.  Readability feels better this way.
	if length < 1 || width < 1 {
		return -1
	}

	// we are assuming width will be less than or equal to
	// length in the code following this block
	if length > width {
		tmp := width
		width = length
		length = tmp
	}

	return length - 1 + length*(width-1)
}

package is_triangle

/*
IsTriangle

Implement a function that accepts 3 integer values a, b, c. The function should return true if a triangle can be
built with the sides of given length and false in any other case.

(In this case, all triangles must have surface greater than 0 to be accepted).

NOTE: not a very good test.  this tests geometry knowledge. not programming
skills.  arguably worse than fizz-buzz.
*/
func IsTriangle(a, b, c int) bool {
	return a+b > c && a+c > b && b+c > a
}

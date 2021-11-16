package string_repeat

import "strings"

/*
RepeatStr

Write a function called repeatStr which repeats the given string exactly n times.

Note: This is a terrible test for most languages.  I could write a loop or some
      sort of builder, but why not use stdlib functionality.
*/
func RepeatStr(repetitions int, value string) string {
	return strings.Repeat(value, repetitions)
}

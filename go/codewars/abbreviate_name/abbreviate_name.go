package abbreviate_name

import (
	"fmt"
	"strings"
)

/*
AbbreviateName

Write a function to convert a name into initials. This kata strictly takes two words with one space in between them.

The output should be two capital letters with a dot separating them.

NOTE: many assumptions here.  Validate input before using in a production environment

NOTE: this code is not optimized.  We are iterating over the length of name multiple times to
      adjust casing and split.  If this were to be called a very, very large amount a single
      loop should be written to improve performance.  But who are we kidding, that will never
      happen.
*/
func AbbreviateName(name string) string {
	words := strings.Split(strings.ToUpper(name), " ")
	// []rune is done to support the full UTF-8 alphabet rather than just ASCII
	return fmt.Sprintf("%s.%s", string([]rune(words[0])[0]), string([]rune(words[1])[0]))
}

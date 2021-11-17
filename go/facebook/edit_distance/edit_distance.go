package edit_distance

import (
	"unicode/utf8"
)

func OneEditApart(s1, s2 string) bool {
	if utf8.RuneCountInString(s1) < utf8.RuneCountInString(s2) {
		tmp := s1
		s1 = s2
		s2 = tmp
	}

	var edits, s1ByteIdx, s2ByteIdx int
	for i := 0; i < utf8.RuneCountInString(s1); i++ {
		s1Rune, s1RuneWidth := utf8.DecodeRuneInString(s1[s1ByteIdx:])
		s1ByteIdx += s1RuneWidth

		s2Rune, s2RuneWidth := utf8.DecodeRuneInString(s2[s2ByteIdx:])
		s2ByteIdx += s2RuneWidth

		if s1Rune == s2Rune {
			continue
		}

		nextS1Rune, _ := utf8.DecodeRuneInString(s1[s1ByteIdx:])
		nextS2Rune, _ := utf8.DecodeRuneInString(s2[s2ByteIdx:])
		if nextS1Rune == nextS2Rune {
			// character replace, no action needed
		} else if nextS1Rune == s2Rune {
			s2ByteIdx -= s2RuneWidth
		} else {
			return false
		}

		edits++
		if edits > 1 {
			return false
		}
	}

	return true
}

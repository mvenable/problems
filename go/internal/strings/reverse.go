package strings

func Reverse(str string) string {
	runes := make([]rune, len(str))

	for i, c := range str {
		runes[len(str)-i-1] = c
	}

	return string(runes)
}

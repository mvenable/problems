package in_array

import (
	"sort"
	"strings"
)

/*
InArray

Given two arrays of strings a1 and a2 return a sorted array r in lexicographical order of the strings of a1 which are
substrings of strings of a2. r must be without duplicates.
*/
func InArray(a1 []string, a2 []string) []string {
	keyMap := map[string]interface{}{}

	// For large lists more optimization could be added.  This is a pure brute-force answer.
	// Loops could be swapped such that a1 was done first.  This change would allow for duplicate
	// checking earlier on.  Left this answer because it is easier to follow and our example inputs
	// were very small.
	for _, superString := range a2 {
		for _, subString := range a1 {
			if strings.Contains(superString, subString) {
				keyMap[subString] = nil
				break
			}
		}
	}

	// Normally I would initialize this as 'var result []string'. However, due to the expected results in the problem
	// we are expecting an empty array response and not a nil value.  :shrug:
	result := []string{}
	for key := range keyMap {
		result = append(result, key)
	}

	sort.Strings(result)
	return result
}

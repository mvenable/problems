package consecutive_strings

import "strings"

/*
LongestConsecutive

You are given an array(list) stringArray of strings and an integer k. Your task is to return the first longest
string consisting of k consecutive strings taken in the array.
*/
func LongestConsecutive(stringArray []string, k int) string {
	longest := ""
	for i := 0; i < len(stringArray)-k+1; i++ {
		current := strings.Join(stringArray[i:i+k], "")
		if len(current) > len(longest) {
			longest = current
		}
	}
	return longest
}

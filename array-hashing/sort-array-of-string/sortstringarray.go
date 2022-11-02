package sortarrayofstring

import "strings"

/*
Sort an array of strings according to string lengths (geeksforgeeks.org/sort-array-strings-according-string-lengths/)
We are given an array of strings, we need to sort the array in increasing order of string lengths.
Examples:

Input : {“GeeksforGeeeks”, “I”, “from”, “am”}
Output : I am from GeeksforGeeks

Input :  {“You”, “are”, “beautiful”, “looking”}
Output : You are looking beautiful
*/

func sortStringArrayBruteForce(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs); i++ {
		for j := i + 1; j < len(strs); j++ {
			if len(strs[i]) > len(strs[j]) {
				//swipe the value
				strs[i], strs[j] = strs[j], strs[i]
			}
		}
	}

	return strings.Join(strs, " ")
}

func sortStringArrayBruteForceImprved(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	for i := 0; i < len(strs); i++ {
		j := i + 1
		for j < len(strs) && len(strs[i]) > len(strs[j]) {
			//swipe the value
			strs[i], strs[j] = strs[j], strs[i]
			j++
		}
	}

	return strings.Join(strs, " ")
}

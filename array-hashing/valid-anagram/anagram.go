package validanagram

/*
242. Valid Anagram - https://leetcode.com/problems/valid-anagram/
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.



Example 1:

Input: s = "anagram", t = "nagaram"
Output: true
Example 2:

Input: s = "rat", t = "car"
Output: false

*/

func isAnagram(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return false
	}

	if len(s) != len(s) {
		return false
	}

	chars := make([]int, 26)

	for _, char := range s {
		chars[char-'a']++
	}
	for _, char := range t {
		chars[char-'a']--
	}

	for i := 0; i < len(chars); i++ {
		if chars[i] > 0 {
			return false
		}
	}

	return true
}

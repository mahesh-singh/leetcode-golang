/*
https://leetcode.com/problems/valid-palindrome/

A phrase is a palindrome if, after converting all uppercase letters into lowercase letters and removing all non-alphanumeric characters, it reads the same forward and backward. Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.



Example 1:

Input: s = "A man, a plan, a canal: Panama"
Output: true
Explanation: "amanaplanacanalpanama" is a palindrome.
Example 2:

Input: s = "race a car"
Output: false
Explanation: "raceacar" is not a palindrome.
Example 3:

Input: s = " "
Output: true
Explanation: s is an empty string "" after removing non-alphanumeric characters.
Since an empty string reads the same forward and backward, it is a palindrome.
*/

package valid_palindrome

import (
	"strings"
	"unicode/utf8"
)

func isPalindrome(s string) bool {

	if len(s) == 0 {
		return true
	}

	left, right := 0, utf8.RuneCountInString(s)-1
	s = strings.ToLower(s)
	for left < right {
		if !isAlphaNum(s[left]) {
			//skip the char
			left = left + 1
		}

		if !isAlphaNum(s[right]) {
			right = right - 1
		}

		if isAlphaNum(s[left]) && isAlphaNum(s[right]) && s[left] != s[right] {
			return false
		}

		if isAlphaNum(s[left]) && isAlphaNum(s[right]) && s[left] == s[right] {
			left = left + 1
			right = right - 1
		}
	}

	return true
}

func isAlphaNum(l byte) bool {
	if (l >= 48 && l <= 57) || (l >= 97 && l <= 122) {
		return true
	}
	return false
}

/*
125. Valid Palindrome https://leetcode.com/problems/valid-palindrome/
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

/*
Notes
- Checking is alpha-numeric, need to create a func based on rune value (0-9 == 48 -57) and (a-z == 97-122)
- Make string to lower by strings.ToLower()
- convert to lower string into []rune
- Two pointer, one start from start called left and another from end called right,
- if left or right pointer char is not a alpha-num, increase/decrease the pointer and contunue the loop
- otherwise check the equality between left pointer char and right pointer char. if they are not equeal return false
*/

package valid_palindrome

import (
	"strings"
)

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	chars := []rune(s)
	left, right := 0, len(chars)-1

	for left < right {
		if !isAlphaNum(s[left]) {
			left++
			continue
		}

		if !isAlphaNum(s[right]) {
			right--
			continue
		}

		if isAlphaNum(s[left]) && isAlphaNum(s[right]) && s[left] != s[right] {
			return false
		}

		left++
		right--
	}

	return true
}

func isAlphaNum(l byte) bool {
	if (l >= 48 && l <= 57) || (l >= 97 && l <= 122) {
		return true
	}
	return false
}

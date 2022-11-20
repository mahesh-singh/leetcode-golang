package validparentheses

import "strings"

/*
20. Valid Parentheses
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

Example 1:

Input: s = "()"
Output: true

Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false
*/

/*
Intuition:
1. Iterating individual chars of input string and using a stack
2. Put opening brackets in stack
3. If closing bracket encounter during iteration, peep the stack to check the corresponding opening bracket exist
4. If exist  - pop, else return false
5. At the end, stack should be empty

Mistakes
#1. Not return just after if closing bracket doesn't match with corresponding opening bracket
#2.
*/

func isValidBruteForce(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	chars := strings.Split(s, "")
	stack := []string{}
	for i := 0; i < len(chars); i++ {
		if chars[i] == "{" || chars[i] == "[" || chars[i] == "(" {
			stack = append(stack, chars[i])
		} else if len(stack) > 0 && (chars[i] == "}" || chars[i] == "]" || chars[i] == ")") {
			l := len(stack) - 1
			if (chars[i] == "}" && stack[l] == "{") || (chars[i] == ")" && stack[l] == "(") || (chars[i] == "]" && stack[l] == "[") {
				stack = stack[0:l]
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return len(stack) == 0
}

/*
Optimization:
1.
*/

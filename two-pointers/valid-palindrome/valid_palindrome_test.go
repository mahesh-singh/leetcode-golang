package valid_palindrome

import "testing"

var validPalindromeTests = []struct {
	in  string
	out bool
}{
	{"A man, a plan, a canal: Panama", true},
	{"race a car", false},
	{" ", true},
}

func TestIsPalindrome(t *testing.T) {
	for _, tt := range validPalindromeTests {
		t.Run(tt.in, func(t *testing.T) {
			got := isPalindrome(tt.in)

			if got != tt.out {
				t.Errorf("got %t, want %t, given %v", got, tt.out, tt.in)
			}
		})
	}
}

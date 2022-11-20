package validparentheses

import "testing"

func TestIsValidParentheses(t *testing.T) {
	cases := []struct {
		input    string
		expected bool
	}{
		{"()",
			true},
		{"()[]{}", true},
		{"(]", false},
		{"([}}])", false},
		{"{[]}", true},
		{"){", false},
		{"([}}])", false},
		{"))", false},
	}

	for _, c := range cases {
		output := isValidBruteForce(c.input)
		if output != c.expected {
			t.Errorf("Expected: %v, output: %v", c.expected, output)
		}
	}
}

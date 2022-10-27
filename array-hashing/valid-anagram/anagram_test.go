package validanagram

import "testing"

func TestValidAnagram(t *testing.T) {
	s, a := "anagram", "nagaram"

	out := isAnagram(s, a)

	if !out {
		t.Errorf("Expected: %v, output: %v", true, out)
	}

	s, a = "aaacc", "cccaa"
	out = isAnagram(s, a)

	if out {
		t.Errorf("Expected: %v, output: %v", false, out)
	}
}

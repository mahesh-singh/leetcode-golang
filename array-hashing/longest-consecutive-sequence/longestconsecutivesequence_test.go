package longestconsecutivesequence

import "testing"

func TestLongestConsecutiveSequenceSorting(t *testing.T) {
	input := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}

	expectedOutput := 7

	output := longestConsecutiveBruteForceSorting(input)

	if output != expectedOutput {
		t.Errorf("Expected: %v, got: %v", expectedOutput, output)
	}
}

func TestLongestconsecutiveSequenceNestedLoop(t *testing.T) {
	input := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}

	expectedOutput := 7

	output := longestConsecutiveBruteForceNestedLoop(input)

	if output != expectedOutput {
		t.Errorf("Expected: %v, got: %v", expectedOutput, output)
	}
}

func TestLongestConsecutiveSequence(t *testing.T) {
	input := []int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}

	expectedOutput := 7

	output := longestConsecutive(input)

	if output != expectedOutput {
		t.Errorf("Expected: %v, got: %v", expectedOutput, output)
	}
}

package sortarrayofstring

import (
	"strings"
	"testing"
)

func TestSortArrayofString(t *testing.T) {
	input := []string{"GeeksforGeeeks", "I", "from", "am"}
	Expectedoutput := "I am from GeeksforGeeeks"

	output := sortStringArrayBruteForce(input)
	if strings.Compare(Expectedoutput, output) != 0 {
		t.Errorf("Expected: %v, output: %v", Expectedoutput, output)
	}

	output = sortStringArrayBruteForceImprved(input)
	if strings.Compare(Expectedoutput, output) != 0 {
		t.Errorf("Expected: %v, output: %v", Expectedoutput, output)
	}
}

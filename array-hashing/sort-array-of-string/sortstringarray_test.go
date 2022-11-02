package sortarrayofstring

import (
	"strings"
	"testing"
)

func TestSortArrayofString(t *testing.T) {
	input := []string{"from", "am", "GeeksforGeeeks", "I"}
	Expectedoutput := "I am from GeeksforGeeeks"

	output := sortStringArrayBruteForce(input)
	if strings.Compare(Expectedoutput, output) != 0 {
		t.Errorf("Expected: %v, output: %v", Expectedoutput, output)
	}

	output2 := sortStringArraycompare(input)
	if strings.Compare(Expectedoutput, output2) != 0 {
		t.Errorf("Expected: %v, output: %v", Expectedoutput, output2)
	}
}

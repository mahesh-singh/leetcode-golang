package groupanagram

import (
	"reflect"
	"testing"
)

func TestGroupAnagram(t *testing.T) {
	input := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	expectedOutput := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}

	out := groupAnagrams(input)
	if !reflect.DeepEqual(expectedOutput, out) {
		t.Errorf("Expected : %v, output: %v", expectedOutput, out)
	}

}

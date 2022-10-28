package groupanagram

import (
	"sort"
	"strings"
)

/*
49. Group Anagrams (https://leetcode.com/problems/group-anagrams/)

Given an array of strings strs, group the anagrams together. You can return the answer in any order.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:

Input: strs = ["eat","tea","tan","ate","nat","bat"]
Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
Example 2:

Input: strs = [""]
Output: [[""]]
Example 3:

Input: strs = ["a"]
Output: [["a"]]
*/

func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{{""}}
	}
	anagramMap := make(map[string][]string)
	var groupAnagram [][]string
	for _, str := range strs {
		s := sortString(str)
		if _, ok := anagramMap[s]; ok {
			anagramMap[s] = append(anagramMap[s], str)
		} else {
			anagramMap[s] = append([]string{}, str)
		}
	}

	for _, v := range anagramMap {
		groupAnagram = append(groupAnagram, v)
	}

	return groupAnagram
}

func sortString(s string) string {
	c := strings.Split(s, "")
	sort.Strings(c)
	return strings.Join(c, "")
}

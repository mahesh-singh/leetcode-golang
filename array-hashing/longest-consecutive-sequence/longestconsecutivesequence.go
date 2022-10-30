package longestconsecutivesequence

import (
	"sort"
)

/*
128. Longest Consecutive Sequence https://leetcode.com/problems/longest-consecutive-sequence/

Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.

You must write an algorithm that runs in O(n) time.



Example 1:

Input: nums = [100,4,200,1,3,2]
Output: 4
Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
Example 2:

Input: nums = [0,3,7,2,5,8,4,6,0,1]
Output: 9
*/

/*
Notes
Brute force
1. Sort the nums array
2. Have two vaiable, one to track the longestStreak and another is currentStreak
3. If currentStreak is greater than longestStreak, update the longestStreak
4. Return the longestStreak
*/

func longestConsecutiveBruteForceSorting(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)

	currentStreak := 1
	longestStreak := 1

	//Notice: len(nums)-1, don't go till end of array. Its upto second last element
	for i := 0; i < len(nums)-1; i++ {
		if nums[i]+1 == nums[i+1] {
			currentStreak++
		} else if nums[i] == nums[i+1] {
			//This case for handling when elements are repeting like [1,1,1,2,4,5]
			continue
		} else {
			//reset the current streak
			currentStreak = 1
		}

		if currentStreak > longestStreak {
			longestStreak = currentStreak
		}
	}
	return longestStreak
}

/*
1. Contains function to check the element in the array
2. For loop on the nums array and for each element check while its next increment exist in the nums or not via above helpher func
3. Incement the currentStreak count
3. If currentStreak is greater than longestStreak, assign currentStreack to longestStreak
*/

func longestConsecutiveBruteForceNestedLoop(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	currentStreak := 1
	longestStreak := 1

	for i := 0; i < len(nums); i++ {
		currentNum := nums[i]
		for containsElements(nums, currentNum+1) {
			currentNum++
			currentStreak++
		}

		if currentStreak > longestStreak {
			longestStreak = currentStreak
		}
		//reset for next iteration
		currentStreak = 1
	}
	return longestStreak
}

func containsElements(nums []int, e int) bool {
	for i := 0; i < len(nums); i++ {
		if nums[i] == e {
			return true
		}
	}
	return false
}

/*
1. Instead of using containstElement func, map can be used which will reduce to O(n^2) from O(n^3)
2. But a trick used in the solution which bring it to O(n)
3. Trick:  - before check for contains;  check if current num-1 not exists before go into the while loop
4. This will make sure to reduce all the duplicate check, means while loop will start from the lowest element
5. It will reduce the complexity to O(n)

Mistakes
1. having main for loop on nums instead of uniqueMap
2. inner loop, not able to formulate the loop over map
3. map structure was map[int]int instead of mapp[int]bool
4. Took time to understand OK := uniqueMap[num-1]
*/

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	uniqueMap := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		uniqueMap[nums[i]] = true
	}

	currentStreack := 0
	longestStreak := 0
	for num := range uniqueMap {
		if _, OK := uniqueMap[num-1]; !OK {

			for currentNum := num; uniqueMap[currentNum]; currentNum++ {
				currentStreack++
			}

			if currentStreack > longestStreak {
				longestStreak = currentStreack
			}
			currentStreack = 0
		}
	}

	return longestStreak
}

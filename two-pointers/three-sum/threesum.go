/*
15. 3Sum
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:

Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.
Example 2:

Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.
Example 3:

Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

*/

package threesum

import "sort"

func threeSum(nums []int) [][]int {
	var result [][]int
	//Sorting is require to remove the duplicate later
	sort.Ints(nums)
	//top Loop will only go though the till len(nums)-2
	for i := 0; i+2 < len(nums); i++ {
		//Skip the duplicate, because array is sorted and current should't be same as previous
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k, target := i+1, len(nums)-1, -nums[i]
		/*
		   Originly I tried to solve via brute force through three innner loop. Not able to remove the duplicate otherwise producing the result
		   To remove one of the loop, take two pointers one (j) start just after the i and another one (k) starts from end.

		*/
		for j < k {
			sum := nums[j] + nums[k]
			if sum == target {

				result = append(result, []int{nums[i], nums[j], nums[k]})
				j++
				k--

				//Skip the duplicate, because array is sorted and current should't be same as previous
				for j < k && nums[j] == nums[j-1] {
					j++
				}
				//Skip the duplicate, because array is sorted and current should't be same as previous
				for j < k && nums[k] == nums[k+1] {
					k--
				}

				//Again sorting help to move the pointer. If sum is greater then target, means move the end pointer (k) to left
				//It will be reduce the sum value, otherwise move the right pointer (j) towards right
			} else if sum > target {
				k--
			} else {
				j++
			}

		}

	}
	return result
}

/*
[-2,0,3,-1,4,0,3,4,1,1,1,-3,-5,4,0]
*/

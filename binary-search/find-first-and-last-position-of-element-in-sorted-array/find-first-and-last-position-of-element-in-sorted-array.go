/*
Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
*/

package find_first_and_last_position_of_element_in_sorted_array

func searchRange(nums []int, target int) []int {
	res := []int{-1, -1}
	if len(nums) == 0 {
		return res
	}

	// find left
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	// don't found target
	if lo >= len(nums) || nums[lo] != target {
		return res
	}

	res[0] = lo

	// find right
	lo, hi = 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	res[1] = lo - 1

	return res
}

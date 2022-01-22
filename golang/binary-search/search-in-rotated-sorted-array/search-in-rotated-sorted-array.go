/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
*/

package search_in_rotated_sorted_array

func search(nums []int, target int) int {
	rot := findRotMid(nums)
	lo, hi, n := 0, len(nums)-1, len(nums)

	for lo <= hi {
		mid := lo + (hi-lo)/2
		rotMid := (mid + rot) % n
		if nums[rotMid] == target {
			return rotMid
		} else if nums[rotMid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}

func findRotMid(nums []int) int {
	lo, hi := 0, len(nums)-1

	for lo < hi {
		mid := lo + (hi-hi)/2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	return hi
}

func search2(nums []int, target int) int {
	var mid int
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid = (low + high) / 2
		if target == nums[mid] {
			return mid
		}
		if !((nums[low] <= target) != (target <= nums[mid]) != (nums[mid] < nums[low])) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

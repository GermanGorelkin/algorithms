/*
Given a sorted (in ascending order) integer array nums of n elements and a target value, write a function to search target in nums. If target exists, then return its index, otherwise return -1.


Example 1:

Input: nums = [-1,0,3,5,9,12], target = 9
Output: 4
Explanation: 9 exists in nums and its index is 4

Example 2:

Input: nums = [-1,0,3,5,9,12], target = 2
Output: -1
Explanation: 2 does not exist in nums so return -1


Note:

You may assume that all elements in nums are unique.
n will be in the range [1, 10000].
The value of each element in nums will be in the range [-9999, 9999].
*/

package binary_search

/*
Distinguishing Syntax:
	Initial Condition: left = 0, right = length-1
	Termination: left > right
	Searching Left: right = mid-1
	Searching Right: left = mid+1
*/
func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return -1
}

/*
It is used to search for an element or condition which requires accessing the current index and
its immediate right neighbor's index in the array.

Distinguishing Syntax:
	Initial Condition: left = 0, right = length
	Termination: left == right
	Searching Left: right = mid
	Searching Right: left = mid+1
*/
func search2(nums []int, target int) int {
	lo, hi := 0, len(nums)

	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	if lo != len(nums) && nums[lo] == target {
		return lo
	}
	return -1
}

/*
It is used to search for an element or condition which requires accessing the current index
and its immediate left and right neighbor's index in the array.

Distinguishing Syntax:
	Initial Condition: left = 0, right = length-1
	Termination: left + 1 == right
	Searching Left: right = mid
	Searching Right: left = mid
*/
func search3(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo+1 < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			hi = mid
		} else {
			lo = mid
		}
	}

	if nums[lo] == target {
		return lo
	}
	if nums[hi] == target {
		return hi
	}
	return -1
}

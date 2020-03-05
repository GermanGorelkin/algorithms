/*
Given an array of n positive integers and a positive integer s, find the minimal length of a contiguous subarray of which the sum ≥ s. If there isn't one, return 0 instead.

Example:

Input: s = 7, nums = [2,3,1,2,4,3]
Output: 2
Explanation: the subarray [4,3] has the minimal length under the problem constraint.
*/

package minimum_size_subarray_sum

func minSubArrayLen(s int, nums []int) int {
	var i, j, sum int
	cnt := int(^uint(0) >> 1) // max int

	for {
		if sum < s && j < len(nums) {
			sum += nums[j]
			j++
		} else if i < len(nums) {
			if sum >= s {
				if cnt > j-i {
					cnt = j - i
				}
				sum -= nums[i]
			}
			i++
		}

		if i >= j {
			break
		}
	}

	return cnt
}

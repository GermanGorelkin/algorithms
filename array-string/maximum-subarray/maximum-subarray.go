/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.
*/

package maximum_subarray

/*
 O(n^3)
 O(1)
*/
func maxSubArrayV1(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}

	max := nums[0]
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			tmp := 0
			for k := i; k <= j; k++ {
				tmp += nums[k]
			}
			if tmp > max {
				max = tmp
			}
		}
	}

	return max
}

/*
 Kadane's algorithm
 O(n)
 O(1)
*/
func maxSubArrayV2(nums []int) int {
	bestSum := nums[0]
	currentSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currentSum = max(currentSum+nums[i], nums[i])
		bestSum = max(bestSum, currentSum)
	}
	return bestSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

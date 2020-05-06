/*
Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.

You may assume that the array is non-empty and the majority element always exist in the array.

Example 1:

Input: [3,2,3]
Output: 3
Example 2:

Input: [2,2,1,1,1,2,2]
Output: 2
*/

package majority_element

func majorityElement(nums []int) int {
	m := make(map[int]int)
	wantCount := len(nums) / 2

	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
		if m[nums[i]] > wantCount {
			return nums[i]
		}
	}
	return -1
}

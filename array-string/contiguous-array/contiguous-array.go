/*
Given a binary array, find the maximum length of a contiguous subarray with equal number of 0 and 1.

Example 1:
Input: [0,1]
Output: 2
Explanation: [0, 1] is the longest contiguous subarray with equal number of 0 and 1.
Example 2:
Input: [0,1,0]
Output: 2
Explanation: [0, 1] (or [1, 0]) is a longest contiguous subarray with equal number of 0 and 1.
Note: The length of the given binary array will not exceed 50,000.
*/

package contiguous_array

func findMaxLength(nums []int) int {
	var maxln int
	var count int
	m := make(map[int]int)

	for i, v := range nums {
		if v == 1 {
			count++
		} else {
			count--
		}
		if count == 0 {
			maxln = i + 1
		} else if prev, ok := m[count]; ok {
			if i-prev > maxln {
				maxln = i - prev
			}
		} else {
			m[count] = i
		}
	}

	return maxln
}

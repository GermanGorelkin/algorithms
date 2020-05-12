/*
You are given a sorted array consisting of only integers where every element appears exactly twice, except for one element which appears exactly once. Find this single element that appears only once.



Example 1:

Input: [1,1,2,3,3,4,4,8,8]
Output: 2
Example 2:

Input: [3,3,7,7,10,11,11]
Output: 10


Note: Your solution should run in O(log n) time and O(1) space.
*/

package single_element_in_a_sorted_array

func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1

	for lo < hi {
		// We want the first element of the middle pair,
		// which should be at an even index if the left part is sorted.
		mid := lo + (hi-lo)/2
		if mid%2 == 1 {
			mid--
		}

		// We didn't find a pair. The single element must be on the left.
		// (pipes mean lo & hi)
		if nums[mid] != nums[mid+1] {
			hi = mid
		} else {
			// We found a pair. The single element must be on the right.
			lo += 2
		}
	}

	// 'hi' should always be at the beginning of a pair.
	// When 'start > end', start must be the single element.
	return nums[hi]
}

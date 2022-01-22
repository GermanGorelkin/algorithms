/*
Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.

Note:

The number of elements initialized in nums1 and nums2 are m and n respectively.
You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
Example:

Input:
nums1 = [1,2,3,0,0,0], m = 3
nums2 = [2,5,6],       n = 3

Output: [1,2,2,3,5,6]

Example 2:

Input:
nums1 = [4,5,6,0,0,0], m = 3
nums2 = [1,2,3],       n = 3

Output: [1,2,3,4,5,6]
*/

package merge_sorted_array

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m + n - 1
	j, k := m-1, n-1
	for j >= 0 && k >= 0 {
		if nums1[j] > nums2[k] {
			nums1[i] = nums1[j]
			j--
		} else {
			nums1[i] = nums2[k]
			k--
		}
		i--
	}
	if k > 0 {
		copy(nums1[:k+1], nums2[:k+1])
	}
	//for k >= 0 {
	//	nums1[i] = nums2[k]
	//	k--
	//	i--
	//}
}

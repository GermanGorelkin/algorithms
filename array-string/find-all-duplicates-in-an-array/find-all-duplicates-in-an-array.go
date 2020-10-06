/*
Given an array of integers, 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.

Find all the elements that appear twice in this array.

Could you do it without extra space and in O(n) runtime?

Example:
Input:
[4,3,2,7,8,2,3,1]

Output:
[2,3]
*/

package find_all_duplicates_in_an_array

func findDuplicates(nums []int) []int {
	res := make([]int, 0)

	for i := range nums {
		j := abs(nums[i])
		if nums[j-1] < 0 {
			res = append(res, j)
		} else {
			nums[j-1] *= -1
		}
	}

	return res
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}

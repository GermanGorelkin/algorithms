package containsduplicateiii

/*
Given an integer array nums and two integers k and t, return true if there are two distinct indices i and j in the array such that abs(nums[i] - nums[j]) <= t and abs(i - j) <= k.
*/

/*
	1,2,3,1


	abs(nums[i] - nums[j]) <= t
	10 - 7 <= 4
	10 <= 7+4
	10 >= 7-4

*/

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {

	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

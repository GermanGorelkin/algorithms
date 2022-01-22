package insertion

func sortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := 1; i < len(nums); i++ {
		for j := i; j > 0 && nums[j] < nums[j-1]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

	return nums
}

package bubble

func sortArray(nums []int) []int {
	swapped := true
	for i := 0; i < len(nums) && swapped; i++ {
		swapped = false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
	}
	return nums
}

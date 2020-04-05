/*
Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.

Example:

Input: [0,1,0,3,12]
Output: [1,3,12,0,0]
Note:

You must do this in-place without making a copy of the array.
Minimize the total number of operations.
*/

package move_zeroes

func moveZeroesV3(nums []int) {
	lastNonZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZero], nums[i] = nums[i], nums[lastNonZero]
			lastNonZero++
		}
	}
}

func moveZeroesV2(nums []int) {
	lastNonZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZero] = nums[i]
			lastNonZero++
		}
	}

	for lastNonZero < len(nums) {
		nums[lastNonZero] = 0
		lastNonZero++
	}
}

func moveZeroes(nums []int) {
	end := len(nums) - 1
	for i := 0; i <= end; {
		if nums[i] == 0 {
			for j := i; j < end; j++ {
				nums[j] = nums[j+1]
			}
			nums[end] = 0
			end--
		} else {
			i++
		}
	}
}

/*
Given an array, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input:  [1,2,3,4,5,6,7] and k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: [-1,-100,3,99] and k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
Note:

Try to come up as many solutions as you can, there are at least 3 different ways to solve this problem.
Could you do it in-place with O(1) extra space?
*/

package rotate_array

/*
	Time complexity : O(n)
	Space complexity : O(n)
*/
func rotate(nums []int, k int) {
	s := make([]int, len(nums))
	copy(s, nums[len(nums)-k:])
	copy(s[k:], nums[:k+1])
	copy(nums, s)
}

/*
	Time complexity : O(n*k)
	Space complexity : O(1)
*/
func rotate2(nums []int, k int) {
	for i := 0; i < k; i++ {
		for j := len(nums) - k + i; j > i; j-- {
			nums[j-1], nums[j] = nums[j], nums[j-1]
		}
	}
}

/*
	Time complexity : O(n)
	Space complexity : O(1)
*/
func rotate3(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}

/*
	Time complexity : O(n)
	Space complexity : O(1)
*/
func rotate4(nums []int, k int) {
	l := len(nums)
	if l <= 1 || k == 0 || k == l {
		return
	}

	next, start := 0, 0
	val := nums[next]

	for i := 0; i < l; i++ {
		next = (next + k) % l
		val, nums[next] = nums[next], val

		if next == start {
			next++
			start = next
			val = nums[next]
		}
	}
}

func rotate5(nums []int, k int) {
	a := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newIndex := (i + k) % len(nums)
		a[newIndex] = nums[i]
	}
	for i := 0; i < len(nums); i++ {
		nums[i] = a[i]
	}
}

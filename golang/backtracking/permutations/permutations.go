/*
Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.



Example 1:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Example 2:

Input: nums = [0,1]
Output: [[0,1],[1,0]]

Example 3:

Input: nums = [1]
Output: [[1]]


Constraints:

1 <= nums.length <= 6
-10 <= nums[i] <= 10
All the integers of nums are unique.
*/

package permutations

/*
[1, 2, 3]

1	1	2	2	3	3
2	3	1	3	1	2
3	2	3	1	2	1

*/

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	backtrack(nums, used, []int{}, &res)
	return res
}

func backtrack(nums []int, used []bool, tmp []int, res *[][]int) {
	if len(nums) == len(tmp) {
		*res = append(*res, append([]int{}, tmp...))
	}

	for i, num := range nums {
		if used[i] {
			continue
		}

		used[i] = true
		tmp = append(tmp, num)
		backtrack(nums, used, tmp, res)
		used[i] = false
		tmp = tmp[:len(tmp)-1]
	}
}

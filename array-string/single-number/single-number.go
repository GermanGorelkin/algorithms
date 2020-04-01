/*
Given a non-empty array of integers, every element appears twice except for one. Find that single one.

Note:

Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

Example 1:

Input: [2,2,1]
Output: 1
Example 2:

Input: [4,1,2,1,2]
Output: 4
*/

package single_number

/*
a^0=a
a^a=0
a^b^a=(a^a)^b=0^b=b
*/
func singleNumber(nums []int) int {
	res := 0
	for _, v := range nums {
		res = res ^ v
	}
	return res
}

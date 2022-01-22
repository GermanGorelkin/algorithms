/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

package plus_one

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	digits[len(digits)-1] += 1

	var n int
	for i := len(digits) - 1; i >= 0; i-- {
		digits[i] += n

		n = digits[i] / 10
		if n == 0 {
			break
		}
		digits[i] %= 10
	}

	if n > 0 {
		digits = append([]int{1}, digits...)
	}

	return digits
}

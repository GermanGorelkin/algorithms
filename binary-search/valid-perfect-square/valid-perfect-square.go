/*
Given a positive integer num, write a function which returns True if num is a perfect square else False.

Note: Do not use any built-in library function such as sqrt.

Example 1:

Input: 16
Output: true
Example 2:

Input: 14
Output: false
*/

package valid_perfect_square

func isPerfectSquare(num int) bool {
	lo, hi := 1, num
	for lo <= hi {
		mid := lo + (hi-lo)/2

		if mid*mid == num {
			return true
		}
		if mid*mid > num {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}

	return false
}

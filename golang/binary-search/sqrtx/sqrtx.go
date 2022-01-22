/*
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2

Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
*/

package sqrtx

func mySqrt(x int) int {
	if x < 0 {
		return 0
	}
	ans, lo, hi := 0, 1, x

	for lo <= hi {
		mid := lo + (hi-lo)/2
		sq := mid * mid
		if sq == x {
			return mid
		} else if sq < x {
			ans = mid
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return ans
}

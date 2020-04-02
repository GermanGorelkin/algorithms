/*
Write an algorithm to determine if a number is "happy".

A happy number is a number defined by the following process: Starting with any positive integer, replace the number by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy numbers.

Example:

Input: 19
Output: true
Explanation:
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
*/

package happy_number

func isHappy(n int) bool {
	m := make(map[int]struct{})

	for {
		digits := split(n)
		n = 0
		for _, v := range digits {
			n += v * v
		}
		if n == 1 {
			return true
		}

		if _, ok := m[n]; ok {
			return false
		}
		m[n] = struct{}{}
	}
}

func split(n int) []int {
	var s []int
	for n != 0 {
		s = append(s, n%10)
		n = n / 10
	}
	return s
}

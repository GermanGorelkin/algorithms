/*
Count the number of prime numbers less than a non-negative number, n.



Example 1:

Input: n = 10
Output: 4
Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
Example 2:

Input: n = 0
Output: 0
Example 3:

Input: n = 1
Output: 0


Constraints:

0 <= n <= 5 * 106
 */

package count_primes

func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	var cnt int

	for n = n - 1; n > 1; n-- {
		prime := true
		for i := n - 1; i > 1; i-- {
			if n%i == 0 {
				prime = false
				break
			}
		}
		if prime {
			cnt++
		}
	}

	return cnt
}

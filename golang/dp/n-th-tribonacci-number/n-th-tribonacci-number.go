/*
	https://leetcode.com/problems/n-th-tribonacci-number/

	The Tribonacci sequence Tn is defined as follows:
	T0 = 0, T1 = 1, T2 = 1, and Tn+3 = Tn + Tn+1 + Tn+2 for n >= 0.
	Given n, return the value of Tn.
*/
package n_th_tribonacci_number

func tribonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}

	t0, t1, t2 := 0, 1, 1

	for i := 3; i <= n; i++ {
		t := t0 + t1 + t2
		t0, t1, t2 = t1, t2, t
	}

	return t2
}

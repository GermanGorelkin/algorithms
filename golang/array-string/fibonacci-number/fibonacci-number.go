/*
The Fibonacci numbers, commonly denoted F(n) form a sequence, called the Fibonacci sequence, such that each number is the sum of the two preceding ones, starting from 0 and 1. That is,

F(0) = 0,   F(1) = 1
F(N) = F(N - 1) + F(N - 2), for N > 1.
Given N, calculate F(N).
*/

package fibonacci_number

func fib(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	f1, f2 := 1, 1
	var f int

	for i := 3; i <= N; i++ {
		f = f1 + f2
		f1 = f2
		f2 = f
	}

	return f
}

func frec(N int) int {
	if N == 1 || N == 2 {
		return 1
	}
	return frec(N-2) + frec(N-1)
}

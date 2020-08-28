package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)

	f := fib(n)

	fmt.Printf("%d", f)
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	f0, f1 := 0, 1

	for i := 2; i <= n; i++ {
		f0, f1 = f1, f0+f1
	}

	return f1
}

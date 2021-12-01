package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	ans := fib(n)
	fmt.Printf("%d", ans)
}

func fib(n int) int {
	f1, f2 := 0, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, (f1+f2)%10
	}

	return f2
}

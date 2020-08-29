package main

import "fmt"

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)

	f := gcd(a, b)

	fmt.Printf("%d", f)
}

func gcd(a, b int) int {
	for b != 0 {
		a %= b
		if a == 0 {
			return b
		}
		b %= a
	}
	return a
}

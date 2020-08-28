package main

import (
	"fmt"
)

func main() {
	var a, b int
	_, _ = fmt.Scanf("%d %d", &a, &b)

	sum := a + b

	fmt.Printf("%d", sum)
}

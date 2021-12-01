package main

import (
	"fmt"
)

/*
	1) multiply_matrices  https://habr.com/ru/post/148336/

	2) Pisano period https://en.wikipedia.org/wiki/Pisano_period
*/

func main() {
	var n, m int
	_, _ = fmt.Scanf("%d %d", &n, &m)
	ans := calc(n, m)
	fmt.Printf("%d", ans%m)
}

func calc(n, m int) int {
	p := pisanoPeriod(uint(m))
	ans := p[n%len(p)]
	return int(ans)
}

func pisanoPeriod(m uint) []uint {
	pp := []uint{0, 1}
	var p, c uint = 0, 1
	for i := uint(2); i <= m*m; i++ {
		p, c = c, (p+c)%m
		pp = append(pp, (pp[i-2]+pp[i-1])%m)
		if pp[i-1] == 0 && pp[i] == 1 {
			break
		}
	}
	return pp[:len(pp)-2]
}

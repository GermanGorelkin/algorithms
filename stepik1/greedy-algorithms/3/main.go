/*
Sample Input 1:
4
Sample Output 1:
2
1 3

Sample Input 2:
6
Sample Output 2:
3
1 2 3
*/

package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	res := maxK(n)

	fmt.Printf("%d\n", len(res))
	for _, k := range res {
		fmt.Printf("%d ", k)
	}
}

func maxK(n int) (res []int) {
	for i := 1; i <= n; i++ {
		next := i + 1
		if n-i-next >= 0 {
			res = append(res, i)
			n -= i
		} else {
			res = append(res, n)
			break
		}
	}

	return res
}

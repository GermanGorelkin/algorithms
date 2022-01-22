package v2

/*
StableCountingSort
 */
func Sort(arr []int) {
	c := make([]int, 6)

	for i := 0; i < len(arr); i++ {
		c[arr[i]]++
	}

	for j := 1; j < len(c); j++ {
		c[j] = c[j] + c[j-1]
	}

	b := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		c[arr[i]] = c[arr[i]] - 1
		b[c[arr[i]]] = arr[i]
	}

	copy(arr, b)
}

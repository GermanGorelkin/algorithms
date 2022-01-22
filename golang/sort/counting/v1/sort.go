package v1

/*
SimpleCountingSort
 */
func Sort(arr []int) {
	c := make([]int, 6)

	for i := 0; i < len(arr); i++ {
		c[arr[i]]++
	}

	var b int
	for j := 0; j < len(c); j++ {
		for i := 0; i < c[j]; i++ {
			arr[b] = j
			b++
		}
	}
}

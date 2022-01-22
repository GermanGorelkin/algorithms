package v2

func QuickSort(data []int, l, r int) {
	if l >= r {
		return
	}

	p := data[l] // pivot = first
	m := partition(data, p, l, r)
	QuickSort(data, l, m-1)
	QuickSort(data, m+1, r)
}

func partition(data []int, p, l, r int) int {
	i := l + 1

	for j := l + 1; j <= r; j++ {
		if data[j] < p {
			data[i], data[j] = data[j], data[i]
			i++
		}
	}

	data[l], data[i-1] = data[i-1], data[l]

	return i - 1
}

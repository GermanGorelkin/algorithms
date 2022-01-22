package merge

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	return merge(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
}

func merge(left []int, right []int) []int {
	lsize, rsize := len(left), len(right)
	res := make([]int, lsize+rsize)
	l, r := 0, 0

	for i := 0; i < lsize+rsize; i++ {
		if l >= lsize {
			copy(res[i:], right[r:])
			return res
		}
		if r >= rsize {
			copy(res[i:], left[l:])
			return res
		}

		if left[l] < right[r] {
			res[i] = left[l]
			l++
		} else {
			res[i] = right[r]
			r++
		}
	}

	return res
}

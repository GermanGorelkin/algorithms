package sort

func Sort(arr []int) {
	size := len(arr)
	var i, j, max int
	for i = 0; i < size-1; i++ {
		max = 0
		for j = 1; j < size-i; j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[size-1-i], arr[max] = arr[max], arr[size-1-i]
	}
}

func Sort2(arr []int) {
	size := len(arr)
	var i, j, min int
	for i = 0; i < size-1; i++ {
		min = i
		for j = i + 1; j < size; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

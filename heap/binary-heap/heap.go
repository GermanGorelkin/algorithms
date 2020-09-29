package binary_heap

/*
children: 2i+1, 2i+2
parent: (i-1)/2

min-binary-heap
parent < children

max-binary-heap
parent > children
*/

type heapType uint8

const (
	minHeap heapType = iota
	maxHeap
)

type heap struct {
	data []int
	t    heapType
}

func (h *heap) SiftUp(i int) {
	if h.t == minHeap {
		minSiftUp(h.data, i)
	} else {
		maxSiftUp(h.data, i)
	}
}

func minSiftUp(data []int, i int) {
	for data[i] < data[(i-1)/2] {
		data[i], data[(i-1)/2] = data[(i-1)/2], data[i]
		i = (i - 1) / 2
	}
}
func maxSiftUp(data []int, i int) {
	for data[i] > data[(i-1)/2] {
		data[i], data[(i-1)/2] = data[(i-1)/2], data[i]
		i = (i - 1) / 2
	}
}

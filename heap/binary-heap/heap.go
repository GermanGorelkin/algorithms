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

func HeapSort(data []int) {
	l := len(data)
	if l == 0 || l == 1 {
		return
	}

	BuildHeap(data, maxHeap)
	size := l - 1

	for i := size; i >= 0; i-- {
		data[0], data[size] = data[size], data[0]
		data = data[:size]
		maxSiftDown(data, 0)
		size--
	}

	data = data[:l-1]
}

func BuildHeap(data []int, t heapType) *heap {
	h := heap{
		data: data,
		t:    t,
	}
	for i := len(h.data) / 2; i >= 0; i-- {
		h.SiftDown(i)
	}
	return &h
}

func (h *heap) ChangePriority(index int, val int) {
	if h.t == minHeap {
		minChangePriority(h.data, index, val)
	} else {
		maxChangePriority(h.data, index, val)
	}
}

func maxChangePriority(data []int, i int, val int) {
	old := data[i]
	data[i] = val

	if val > old {
		maxSiftUp(data, i)
	} else {
		maxSiftDown(data, i)
	}
}
func minChangePriority(data []int, i int, val int) {
	old := data[i]
	data[i] = val

	if val > old {
		minSiftDown(data, i)
	} else {
		minSiftUp(data, i)
	}
}

func (h *heap) SiftDown(i int) {
	if h.t == minHeap {
		minSiftDown(h.data, i)
	} else {
		maxSiftDown(h.data, i)
	}
}

/*
func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
*/
/*
	                  9
	               /     \
		          6       4
		        /   \   /   \
		       9     8 12    7
		      / \
		     11

		    0 1 2 3 4 5  6  7 8
		    3 6 4 9 8 12 7 11 9
*/
func minSiftDown(data []int, i int) {
	length := len(data)
	for 2*i+1 < length {
		left := 2*i + 1
		right := 2*i + 2

		smallest := left
		if right < length && data[right] < data[smallest] {
			smallest = right
		}

		if data[i] <= data[smallest] {
			break
		}

		data[smallest], data[i] = data[i], data[smallest]
		i = smallest
	}
}
func maxSiftDown(data []int, i int) {
	length := len(data)
	for 2*i+1 < length {
		left := 2*i + 1
		right := 2*i + 2

		largest := left
		if right < length && data[right] > data[largest] {
			largest = right
		}

		if data[i] >= data[largest] {
			break
		}

		data[largest], data[i] = data[i], data[largest]
		i = largest
	}
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

func (h *heap) Insert(val int) {
	h.data = append(h.data, val)
	h.SiftUp(len(h.data) - 1)
}

func (h *heap) Extract() int {
	if len(h.data) == 0 {
		return 0
	}

	val := h.data[0]
	h.data[0] = h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	h.SiftDown(0)

	return val
}

func (h *heap) Remove(i int) {
	l := len(h.data)
	if l == 0 || l < i {
		return
	}
	if l == 1 {
		h.data = h.data[:0]
		return
	}
	if l-1 == i {
		h.data = h.data[:len(h.data)-1]
		return
	}

	val := h.data[l-1]
	h.data = h.data[:l-1]
	h.ChangePriority(i, val)
}

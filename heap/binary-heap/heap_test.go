package binary_heap

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Insert(t *testing.T) {
	t.Run("min", func(t *testing.T) {
		h := heap{t: minHeap}

		h.Insert(10)
		assert.Equal(t, []int{10}, h.data)

		h.Insert(9)
		assert.Equal(t, []int{9, 10}, h.data)

		h.Insert(8)
		assert.Equal(t, []int{8, 10, 9}, h.data)

		h.Insert(7)
		assert.Equal(t, []int{7, 8, 9, 10}, h.data)

		h.Insert(11)
		assert.Equal(t, []int{7, 8, 9, 10, 11}, h.data)
	})
	t.Run("max", func(t *testing.T) {
		h := heap{t: maxHeap}

		h.Insert(10)
		assert.Equal(t, []int{10}, h.data)

		h.Insert(9)
		assert.Equal(t, []int{10, 9}, h.data)

		h.Insert(11)
		assert.Equal(t, []int{11, 9, 10}, h.data)

		h.Insert(20)
		assert.Equal(t, []int{20, 11, 10, 9}, h.data)

		h.Insert(200)
		assert.Equal(t, []int{200, 20, 10, 9, 11}, h.data)
	})
}

func Test_Extract(t *testing.T) {
	t.Run("min", func(t *testing.T) {
		h := heap{
			data: []int{3, 6, 4, 9, 8, 12, 7, 11, 9},
			t:    minHeap,
		}

		data := make([]int, len(h.data))
		copy(data, h.data)
		sort.Ints(data)

		for _, v := range data {
			assert.Equal(t, v, h.Extract())
		}
	})
	t.Run("max", func(t *testing.T) {
		h := heap{
			data: []int{100, 60, 40, 35, 8, 12, 7, 11, 9},
			t:    maxHeap,
		}

		data := make([]int, len(h.data))
		copy(data, h.data)
		sort.Slice(data, func(i, j int) bool {
			return data[i] > data[j]
		})

		for _, v := range data {
			assert.Equal(t, v, h.Extract())
		}
	})
}

func Test_minSiftDown(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		data := []int{3}
		want := []int{3}

		minSiftDown(data, 0)

		assert.Equal(t, want, data)
	})

	t.Run("1", func(t *testing.T) {
		data := []int{3, 4}
		want := []int{3, 4}

		minSiftDown(data, 0)

		assert.Equal(t, want, data)
	})
	t.Run("2", func(t *testing.T) {
		data := []int{4, 3}
		want := []int{3, 4}

		minSiftDown(data, 0)

		assert.Equal(t, want, data)
	})

	t.Run("3", func(t *testing.T) {
		data := []int{3, 4, 2}
		want := []int{2, 4, 3}

		minSiftDown(data, 0)

		assert.Equal(t, want, data)
	})

	/*
		                  100
		               /     \
			          6       4
			        /   \   /   \
			       9     8 12    7
			      / \
			     11  9

			    0 1 2 3 4 5  6  7 8
			    3 6 4 9 8 12 7 11 9
	*/
	t.Run("4", func(t *testing.T) {
		data := []int{100, 6, 4, 9, 8, 12, 7, 11, 9}
		// 4, 6, 100, 9, 8, 12, 7, 11, 9
		// 4, 6, 7, 9, 8, 12, 100, 11, 9
		want := []int{4, 6, 7, 9, 8, 12, 100, 11, 9}

		minSiftDown(data, 0)

		assert.Equal(t, want, data)
	})
}

func Test_maxSiftDown(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		data := []int{3}
		want := []int{3}

		maxSiftDown(data, 0)

		assert.Equal(t, want, data)
	})

	t.Run("1", func(t *testing.T) {
		data := []int{4, 3}
		want := []int{4, 3}

		maxSiftDown(data, 0)

		assert.Equal(t, want, data)
	})
	t.Run("2", func(t *testing.T) {
		data := []int{3, 4}
		want := []int{4, 3}

		maxSiftDown(data, 0)

		assert.Equal(t, want, data)
	})

	t.Run("3", func(t *testing.T) {
		data := []int{2, 3, 4}
		want := []int{4, 3, 2}

		maxSiftDown(data, 0)

		assert.Equal(t, want, data)
	})

	/*
		                     1
			               /     \
				          60      40
				        /   \   /   \
				       35    8 12    7
				      / \
				     11  9

				    1 60 40 35 8 12 7 11 9
	*/
	t.Run("4", func(t *testing.T) {
		data := []int{1, 60, 40, 35, 8, 12, 7, 11, 9}
		// 60 1 40 35 8 12 7 11 9
		// 60 35 40 1 8 12 7 11 9
		// 60 35 40 11 8 12 7 1 9

		want := []int{60, 35, 40, 11, 8, 12, 7, 1, 9}

		maxSiftDown(data, 0)

		assert.Equal(t, want, data)
	})
}

func Test_minSiftUp(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		data := []int{3}
		want := []int{3}

		minSiftUp(data, 0)

		assert.Equal(t, want, data)
	})

	t.Run("1", func(t *testing.T) {
		data := []int{3, 4}
		want := []int{3, 4}

		minSiftUp(data, 1)

		assert.Equal(t, want, data)
	})
	t.Run("2", func(t *testing.T) {
		data := []int{4, 3}
		want := []int{3, 4}

		minSiftUp(data, 1)

		assert.Equal(t, want, data)
	})

	t.Run("3", func(t *testing.T) {
		data := []int{3, 4, 2}
		want := []int{2, 4, 3}

		minSiftUp(data, 2)

		assert.Equal(t, want, data)
	})

	/*
		                  3
		               /     \
			          6       4
			        /   \   /   \
			       9     8 12    7
			      / \
			     11  9

			    0 1 2 3 4 5  6  7 8
			    3 6 4 9 8 12 7 11 9
	*/
	t.Run("4", func(t *testing.T) {
		data := []int{6, 4, 9, 8, 12, 7, 11, 9, 3}
		want := []int{3, 6, 9, 4, 12, 7, 11, 9, 8}

		minSiftUp(data, 8)

		assert.Equal(t, want, data)
	})
}

func Test_maxSiftUp(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		data := []int{3}
		want := []int{3}

		maxSiftUp(data, 0)

		assert.Equal(t, want, data)
	})

	t.Run("1", func(t *testing.T) {
		data := []int{4, 3}
		want := []int{4, 3}

		maxSiftUp(data, 1)

		assert.Equal(t, want, data)
	})
	t.Run("2", func(t *testing.T) {
		data := []int{3, 4}
		want := []int{4, 3}

		maxSiftUp(data, 1)

		assert.Equal(t, want, data)
	})

	t.Run("3", func(t *testing.T) {
		data := []int{2, 3, 4}
		want := []int{4, 3, 2}

		maxSiftUp(data, 2)

		assert.Equal(t, want, data)
	})

	/*
		                     100
			               /     \
				          60      40
				        /   \   /   \
				       35    8 12    7
				      / \
				     11  9

				    100 60 40 35 8 12 7 11 9
	*/
	t.Run("4", func(t *testing.T) {
		data := []int{60, 40, 35, 8, 12, 7, 11, 9, 100}
		// 60, 40, 35, 100, 12, 7, 11, 9, 8
		// 60, 100, 35, 40, 12, 7, 11, 9, 8
		// 100, 60, 35, 40, 12, 7, 11, 9, 8

		want := []int{100, 60, 35, 40, 12, 7, 11, 9, 8}

		maxSiftUp(data, 8)

		assert.Equal(t, want, data)
	})
}

func Test_minChangePriority(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		data := []int{3}
		want := []int{0}

		minChangePriority(data, 0, 0)

		assert.Equal(t, want, data)
	})

	t.Run("1", func(t *testing.T) {
		data := []int{3, 4}
		want := []int{0, 3}

		minChangePriority(data, 1, 0)

		assert.Equal(t, want, data)
	})
	t.Run("2", func(t *testing.T) {
		data := []int{3, 4}
		want := []int{0, 4}

		minChangePriority(data, 0, 0)

		assert.Equal(t, want, data)
	})

	t.Run("3", func(t *testing.T) {
		data := []int{3, 4, 2}
		want := []int{1, 3, 2}

		minChangePriority(data, 1, 1)

		assert.Equal(t, want, data)
	})

	/*
		                  3
		               /     \
			          6       4
			        /   \   /   \
			       9     8 12    7
			      / \
			     11  9

			    0 1 2 3 4 5  6  7 8

	*/
	t.Run("4", func(t *testing.T) {
		data := []int{6, 4, 9, 8, 12, 7, 11, 9, 3}
		want := []int{1, 6, 9, 8, 4, 7, 11, 9, 3}

		minChangePriority(data, 4, 1)

		assert.Equal(t, want, data)
	})
}

func Test_maxChangePriority(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		data := []int{3}
		want := []int{0}

		maxChangePriority(data, 0, 0)

		assert.Equal(t, want, data)
	})

	t.Run("1", func(t *testing.T) {
		data := []int{4, 3}
		want := []int{3, 0}

		maxChangePriority(data, 0, 0)

		assert.Equal(t, want, data)
	})
	t.Run("2", func(t *testing.T) {
		data := []int{4, 3}
		want := []int{5, 4}

		maxChangePriority(data, 1, 5)

		assert.Equal(t, want, data)
	})

	t.Run("3", func(t *testing.T) {
		data := []int{2, 3, 4}
		want := []int{5, 3, 2}

		maxChangePriority(data, 2, 5)

		assert.Equal(t, want, data)
	})

	/*
		                     100
			               /     \
				          60      40
				        /   \   /   \
				       35    8 12    7
				      / \
				     11  9

				    100 60 40 35 8 12 7 11 9
	*/
	t.Run("4", func(t *testing.T) {
		data := []int{100, 60, 40, 35, 8, 12, 7, 11, 9}
		want := []int{350, 100, 40, 60, 8, 12, 7, 11, 9}

		maxChangePriority(data, 3, 350)

		assert.Equal(t, want, data)
	})
}

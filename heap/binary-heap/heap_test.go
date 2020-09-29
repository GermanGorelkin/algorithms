package binary_heap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

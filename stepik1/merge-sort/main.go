/*
https://stepik.org/lesson/13248/step/5?unit=3433
Задача на программирование: число инверсий
*/
package main

func sort(data []int) []int {
	l := len(data)
	if l <= 1 {
		return data
	}
	return merge(sort(data[:l/2]), sort(data[l/2:]))
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

func sortIter(data []int) []int {
	q := queue{}
	for _, v := range data {
		q.push([]int{v})
	}

	for len(q.data) > 1 {
		q.push(merge(q.pop(), q.pop()))
	}

	return q.pop()
}

type queue struct {
	data [][]int
}

func (q *queue) push(v []int) {
	q.data = append(q.data, v)
}
func (q *queue) pop() []int {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}

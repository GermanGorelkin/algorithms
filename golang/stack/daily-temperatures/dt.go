package dt

import (
	"math"
)

type stack struct {
	s []int
}

func (s *stack) push(v int) {
	s.s = append(s.s, v)
}
func (s *stack) pop() int {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}
func (s *stack) top() int {
	return s.s[len(s.s)-1]
}
func (s *stack) isEmpty() bool {
	return len(s.s) == 0
}

/*
7|73|  |0
6|76|7(73)|0
5|72|6(76)|6-5=1
4|69|5(72),6(76)|5-4=1
3|71|4(69),5(72),6(76)|5-3=2
2|75|3(71),5(72),6(76)|6-2=4
1|74|2(75)6(76)|2-1=1
0|73|1(74),6(76)|1-0=1
*/
func dailyTemperatures3(t []int) []int {
	s := new(stack)
	res := make([]int, len(t))

	for i := 0; i < len(t); i++ {
		for !s.isEmpty() && t[i] > t[s.top()] {
			lowerIndex := s.pop()
			res[lowerIndex] = i - lowerIndex
		}
		s.push(i)
	}

	return res
}

func dailyTemperatures2(T []int) []int {
	nextWarmerIndex := make([]int, 101)
	daysUntilWarmer := make([]int, len(T))
	for i := len(T) - 1; i >= 0; i-- {
		temperature := T[i]
		daysUntilWarmer[i] = int(math.Max(float64(nextWarmerIndex[temperature]-i), 0))

		for j := temperature - 1; j >= 30; j-- {
			nextWarmerIndex[j] = i
		}
	}

	return daysUntilWarmer
}

func dailyTemperatures(t []int) []int {
	maxIndex := len(t)

	next := make([]int, 101)
	for i := range next {
		next[i] = maxIndex
	}

	res := make([]int, len(t))

	for i := len(t) - 1; i >= 0; i-- {
		warmerIndex := maxIndex

		for j := t[i] + 1; j <= 100; j++ {
			if next[j] < warmerIndex {
				warmerIndex = next[j]
			}
		}
		if warmerIndex < maxIndex {
			res[i] = warmerIndex - i
		}
		next[t[i]] = i
	}

	return res
}

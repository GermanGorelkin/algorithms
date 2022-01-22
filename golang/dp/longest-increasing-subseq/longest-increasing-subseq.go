package longest_increasing_subseq

import (
	"log"
	"math"
)

// dp O(n^2)
// кол-во
func LIS1(arr []int) int {
	dp := make([]int, len(arr))
	var result int

	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}

	return result
}

// dp O(n^2)
// восстановления
func LIS2(arr []int) []int {
	dp := make([]int, len(arr))
	var maxIndex int
	var maxVal int

	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		if maxVal < dp[i] {
			maxVal = dp[i]
			maxIndex = i
		}
	}

	result := make([]int, maxVal)

	for i := len(result) - 1; i >= 0; i-- {
		result[i] = arr[maxIndex]
		for j := maxIndex - 1; j >= 0; j-- {
			if dp[j] == dp[maxIndex]-1 && arr[j] < arr[maxIndex] {
				maxIndex = j
				break
			}
		}
	}

	return result
}

// dp O(n^2)
// восстановления с массивом предков
func LIS3(arr []int) []int {
	dp := make([]int, len(arr))
	p := make([]int, len(arr))
	var maxIndex int
	var maxVal int

	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i] > arr[j] {
				dp[i] = max(dp[i], dp[j]+1)
				p[i] = j
			}
		}
		if maxVal < dp[i] {
			maxVal = dp[i]
			maxIndex = i
		}
	}

	result := make([]int, maxVal)

	for i := len(result) - 1; i >= 0; i-- {
		result[i] = arr[maxIndex]
		maxIndex = p[maxIndex]
	}

	return result
}

// dp O(n^2)
// dp[i] это число, на которое оканчивается возрастающая подпоследовательность длины i
// (а если таких чисел несколько — то наименьшее из них).
/*0    1|  2| 3| 4 | 5| 6| 7  |  8|

	  10|  9| 2| 5 | 3| 7| 101| 1|
 10   10|   |  |   |  |  |    |   |
  9   9 |   |  |   |  |  |    |   |
  2   2 |   |  |   |  |  |    |   |
  5   2 |  5|  |   |  |  |    |   |
  3   2 |  3|  |   |  |  |    |   |
  7	  2 |  3| 7|   |  |  |    |   |
 101  2 |  3| 7|101|  |  |    |   |
  1   1 |  3| 7|101|  |  |    |   |

	  8    5  6   7
*/
func LIS4(arr []int) int {
	dp := make([]int, len(arr)+1)
	dp[0] = math.MinInt64
	for i := 1; i <= len(arr); i++ {
		dp[i] = math.MaxInt64
	}
	var result int

	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(dp); j++ {
			if dp[j-1] < arr[i] && arr[i] < dp[j] {
				dp[j] = arr[i]
				result = max(result, j)
			}
		}
	}

	// var result int
	// for result = len(dp) - 1; result >= 0 && dp[result] == math.MaxInt64; result-- {
	// }

	return result
}

/*
O(nlogn)
*/
func LIS5(arr []int) int {
	dp := make([]int, len(arr)+1)
	dp[0] = math.MinInt64
	for i := 1; i <= len(arr); i++ {
		dp[i] = math.MaxInt64
	}

	var result int

	for i := 0; i < len(arr); i++ {
		j := upper_bound(dp, arr[i])
		if dp[j-1] < arr[i] {
			dp[j] = arr[i]
			result = max(result, j)
		}
	}

	return result
}

/*
O(nlogn)
восстановления
*/
func LIS6(arr []int) []int {
	dp := make([]int, len(arr)+1)
	dp[0] = math.MinInt64
	for i := 1; i <= len(arr); i++ {
		dp[i] = math.MaxInt64
	}
	p := make([]int, len(arr))

	var length int

	for i := 0; i < len(arr); i++ {
		j := upper_bound(dp, arr[i])
		if dp[j-1] < arr[i] {
			dp[j] = arr[i]

			p[i] = dp[j-1]

			length = max(length, j)
		}
	}

	result := make([]int, length)

	log.Println(p)
	log.Println(dp)
	// for i := len(result) - 1; i >= 0; i-- {
	// 	result[i] = arr[p[length]]
	// 	length--
	// }

	return result
}

/*
O(nlogn)
восстановления
*/
func LDS1(arr []int) []int {
	dp := make([]int, len(arr)+1)
	dp[0] = math.MaxInt64
	for i := 1; i <= len(arr); i++ {
		dp[i] = math.MinInt64
	}
	p := make([]int, len(arr))

	var length int

	for i := 0; i < len(arr); i++ {
		j := lower_bound(dp, arr[i])
		if dp[j-1] > arr[i] {
			dp[j] = arr[i]
			p[j] = i
			length = max(length, j)
		}
	}

	log.Println(dp)
	result := make([]int, length)

	for i := 0; i < len(result); i++ {
		result[i] = arr[p[length]]
		length--
	}

	return result
}

func ceilIndex(nums []int, increasingSub []int, end int, currIdx int) int {
	low := 1
	high := end
	for low <= high {
		mid := low + (high-low)/2
		if nums[increasingSub[mid]] < nums[currIdx] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func LIS7(nums []int) []int {
	parent := make([]int, len(nums))          // Tracking the predecessors/parents of elements of each subsequence.
	increasingSub := make([]int, len(nums)+1) // Tracking ends of each increasing subsequence.
	length := 0                               // Length of longest subsequence.

	for i := 0; i < len(nums); i++ {
		// Binary search
		low := ceilIndex(nums, increasingSub, length, i)
		pos := low
		// update parent/previous element for LIS
		parent[i] = increasingSub[pos-1]
		// Replace or append
		increasingSub[pos] = i

		// Update the length of the longest subsequence.
		if pos > length {
			length = pos
		}
	}

	// Generate LIS by traversing parent array
	LIS := make([]int, length)
	k := increasingSub[length]
	for j := length - 1; j >= 0; j-- {
		LIS[j] = nums[k]
		k = parent[k]
	}

	return LIS
}

/*
возвращает позицию первого элемента, строго большего данного
*/
func upper_bound(arr []int, x int) int {
	l, r := 0, len(arr)-1

	for l < r {
		m := l + (r-l)/2
		if arr[m] > x {
			r = m
		} else {
			l = m + 1
		}
	}

	return r
}

/*
arr = 3 2 1
x   = 3
ans = 1
*/
func lower_bound(arr []int, x int) int {
	l, r := 0, len(arr)-1

	for l < r {
		m := l + (r-l)/2
		if arr[m] < x {
			r = m
		} else {
			l = m + 1
		}
	}

	return r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

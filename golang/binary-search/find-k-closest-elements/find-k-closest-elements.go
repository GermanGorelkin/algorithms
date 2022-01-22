/*
Given a sorted array arr, two integers k and x, find the k closest elements to x in the array. The result should also be sorted in ascending order. If there is a tie, the smaller elements are always preferred.



Example 1:

Input: arr = [1,2,3,4,5], k = 4, x = 3
Output: [1,2,3,4]
Example 2:

Input: arr = [1,2,3,4,5], k = 4, x = -1
Output: [1,2,3,4]


Constraints:

1 <= k <= arr.length
1 <= arr.length <= 10^4
Absolute value of elements in the array and x will not exceed 104

*/

package find_k_closest_elements

func findClosestElements(arr []int, k int, x int) []int {
	if len(arr) <= k {
		return arr
	}
	if arr[0] >= x {
		return arr[:k]
	}
	if arr[len(arr)-1] <= x {
		return arr[len(arr)-k:]
	}

	var mid int
	lo, hi := 0, len(arr)-1
	for lo <= hi {
		mid = lo + (hi-lo)/2
		if arr[mid] == x {
			break
		}
		if arr[mid] > x {
			hi = mid - 1
		} else {
			lo = mid + 1
		}
	}
	begin, end := max(0, mid-k-1), min(len(arr)-1, mid+k+1)

	for end-begin+1 != k {
		if x-arr[begin] <= arr[end]-x {
			end--
		} else {
			begin++
		}
	}

	return arr[begin : end+1]
}

//
//func findClosestElements(arr []int, k int, x int) []int {
//	if len(arr) <= k {
//		return arr
//	}
//	if arr[0] >= x {
//		return arr[:k]
//	}
//	if arr[len(arr)-1] <= x {
//		return arr[len(arr)-k:]
//	}
//
//	var mid int
//	lo, hi := 0, len(arr)-1
//	for lo+1 < hi {
//		mid = lo + (hi-lo)/2
//		if arr[mid] == x || (arr[mid-1] > x && arr[mid+1] > x) {
//			break
//		}
//		if arr[mid] > x {
//			hi = mid
//		} else {
//			lo = mid
//		}
//	}
//	if arr[lo] == x {
//		mid = lo
//	} else if arr[hi] == x {
//		mid = hi
//	}
//
//	begin, end := max(0, mid-k-1), min(len(arr)-1, mid+k+1)
//
//	for end-begin+1 != k {
//		if x-arr[begin] <= arr[end]-x {
//			end--
//		} else {
//			begin++
//		}
//	}
//
//	return arr[begin : end+1]
//}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func absDiff(a, b int) int {
	x := a - b
	if x < 0 {
		return x * -1
	}
	return x
}

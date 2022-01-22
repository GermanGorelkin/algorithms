/*
Given an array of integers A sorted in non-decreasing order, return an array of the squares of each number, also in sorted non-decreasing order.



Example 1:

Input: [-4,-1,0,3,10]
Output: [0,1,9,16,100]
Example 2:

Input: [-7,-3,2,3,11]
Output: [4,9,9,49,121]


Note:

1 <= A.length <= 10000
-10000 <= A[i] <= 10000
A is sorted in non-decreasing order.
 */

package squares_of_a_sorted_array

func sortedSquares(A []int) []int {
	l := len(A)
	res := make([]int, l, 0)

	var mid int
	for ; mid < l && A[mid] < 0; mid++ {
	}

	i, j := 0, mid
	for i < mid && j < l {
		if A[i] <= A[j] {
			res = append(res, A[i])
			i++
		} else {
			res = append(res, A[i])
			i++
		}
	}

	for ; i < mid; i++ {
		res = append(res, A[i])
	}
	for ; j < l; j++ {
		res = append(res, A[j])
	}

	return res
}

/*
Given a matrix of m x n elements (m rows, n columns), return all elements of the matrix in spiral order.

Example 1:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
Output: [1,2,3,6,9,8,7,4,5]
Example 2:

Input:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
Output: [1,2,3,4,8,12,11,10,9,5,6,7]
*/

package spiral_matrix

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	var res []int

	rowBegin, rowEnd := 0, len(matrix)-1
	colBegin, colEnd := 0, len(matrix[0])-1

	for rowBegin <= rowEnd && colBegin <= colEnd {
		for i := colBegin; i <= colEnd; i++ {
			res = append(res, matrix[rowBegin][i])
		}

		for i := rowBegin + 1; i <= rowEnd; i++ {
			res = append(res, matrix[i][colEnd])
		}

		if rowBegin < rowEnd && colBegin < colEnd {
			for i := colEnd - 1; i > colBegin; i-- {
				res = append(res, matrix[rowEnd][i])

			}
			for i := rowEnd; i > rowBegin; i-- {
				res = append(res, matrix[i][colBegin])
			}
		}

		rowBegin++
		rowEnd--
		colBegin++
		colEnd--
	}

	return res
}

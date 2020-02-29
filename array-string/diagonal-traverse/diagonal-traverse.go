package diagonal_traverse

/*
Given a matrix of M x N elements (M rows, N columns),
return all elements of the matrix in diagonal order as shown in the below image.



Example:

Input:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]

Output:  [1,2,4,7,5,3,6,8,9]

Explanation:



Note:

The total number of elements of the given matrix will not exceed 10,000.
*/

/*
[0,0],[0,1],[0,2],
[1,0],[1,1],[1,2],
[2,0],[2,1],[2,2],

[0,0], [0,1] [1,0], [2,0] [1,1] [0,2], [1,2] [2,1], [2,2]
*/

/*
Time Complexity: O(rows*cols)
Space Complexity: O(min(rows,cols))
*/
func findDiagonalOrder(matrix [][]int) []int {
	rows := len(matrix)
	if rows == 0 {
		return nil
	}

	cols := len(matrix[0])
	var res []int
	var s []int
	var row, col int

	for i := 0; i < rows+cols-1; i++ {
		s = s[:0]

		if i < cols {
			row = 0
		} else {
			row = i - cols + 1
		}

		if i < cols {
			col = i
		} else {
			col = cols - 1
		}

		for row < rows && col > -1 {
			s = append(s, matrix[row][col])
			row++
			col--
		}

		if i%2 == 0 {
			reverse(s)
		}
		res = append(res, s...)
	}

	return res
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

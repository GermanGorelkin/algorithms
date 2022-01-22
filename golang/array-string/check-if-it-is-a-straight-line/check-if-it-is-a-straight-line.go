/*
You are given an array coordinates, coordinates[i] = [x, y], where [x, y] represents the coordinate of a point. Check if these points make a straight line in the XY plane.





Example 1:



Input: coordinates = [[1,2],[2,3],[3,4],[4,5],[5,6],[6,7]]
Output: true
Example 2:



Input: coordinates = [[1,1],[2,2],[3,4],[4,5],[5,6],[7,7]]
Output: false


Constraints:

2 <= coordinates.length <= 1000
coordinates[i].length == 2
-10^4 <= coordinates[i][0], coordinates[i][1] <= 10^4
coordinates contains no duplicate point.
*/

package check_if_it_is_a_straight_line

// find slope from point 1 and point 2
// compare the slopes of all other points wrt to the original slope
//
// Slope of points (x1, y1) and (x2, y2) is:
// (y2 - y1) / (x2 - x1)
//
// Slope of points (x1, y1) and (x3, y3) is:
// (y3 - y1) / (x3 - x1)
//
// for all three points to be on the same line, the slopes should be equal, in other words:
// (x3 - x1)*(y2 - y1) = (y3 - y1)*(x2 - x1)
//

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}
	x1, y1, x2, y2 := coordinates[0][0], coordinates[0][1], coordinates[1][0], coordinates[1][1]

	for i := 2; i < len(coordinates); i++ {
		if (coordinates[i][0]-x1)*(y2-y1) != (coordinates[i][1]-y1)*(x2-x1) {
			return false
		}
	}
	return true
}

package minimum_time_visiting_all_points

func minTimeToVisitAllPoints(points [][]int) int {
	var time int
	
	for i:=1;i<len(points);i++{
		diffX := abs(points[i][0]-points[i-1][0])
		diffY := abs(points[i][1]-points[i-1][1])

		time += max(diffX, diffY)
	}

    return time
}

func abs(a int) int{
	if a < 0 {
		return -a
	}
	return a
}

func max(a,b int) int {
	if a>b{
		return a
	}
	return b
}
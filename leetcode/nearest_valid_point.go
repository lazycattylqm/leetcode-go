package leetcode

import "math"

func nearestValidPoint(x int, y int, points [][]int) int {
	minDist := math.MaxInt32
	minIndex := -1

	for i, point := range points {
		if valid(x, y, point) {
			dist := distance(x, y, point)
			if dist < minDist {
				minDist = dist
				minIndex = i
			}
		}
	}

	return minIndex

}

func valid(x int, y int, point []int) bool {
	return point[0] == x || point[1] == y
}

func distance(x int, y int, point []int) int {
	return abs(x-point[0]) + abs(y-point[1])
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

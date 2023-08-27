package leetcode

import "sort"

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var res [][]int
	current := intervals[0]
	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		if next[0] <= current[1] {
			current[1] = max(current[1], next[1])
		} else {
			res = append(res, current)
			current = next
		}
	}
	res = append(res, current)

	return res

}

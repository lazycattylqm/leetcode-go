package leetcode

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {

	res := make([][]int, 0)
	temp := make([][]int, 0)
	temp = append(temp, newInterval)
	temp = append(temp, intervals...)
	sort.Slice(temp, func(i, j int) bool {
		return temp[i][0] < temp[j][0]

	})

	current := temp[0]

	for i := 1; i < len(temp); i++ {
		next := temp[i]
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

func Insert(intervals [][]int, newInterval []int) [][]int {
	return insert(intervals, newInterval)
}

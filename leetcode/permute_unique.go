package leetcode

import "sort"

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	length := len(nums)
	visited := make([]bool, length)
	path := make([]int, 0)
	sort.Ints(nums)
	var backtracking func(index int)
	backtracking = func(index int) {
		if index == length {
			result = append(result, append([]int{}, path...))
			return
		}
		for i, num := range nums {
			if visited[i] {
				continue
			}
			if i > 0 && num == nums[i-1] && !visited[i-1] {
				continue
			}
			visited[i] = true
			path = append(path, num)
			backtracking(index + 1)
			path = path[:len(path)-1]
			visited[i] = false
		}

	}
	backtracking(0)
	return result
}

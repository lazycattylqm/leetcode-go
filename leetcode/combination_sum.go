package leetcode

func combinationSum3(k int, n int) [][]int {
	length := k
	target := n
	path := make([]int, 0)
	result := make([][]int, 0)
	var calc func(int, int)
	calc = func(now int, sum int) {
		if len(path) == length {
			if sum == target {
				result = append(result, append([]int{}, path...))
			}
			return
		}
		for i := now; i <= 9; i++ {
			path = append(path, i)
			calc(i+1, sum+i)
			path = path[:len(path)-1]
		}
	}
	calc(1, 0)
	return result
}

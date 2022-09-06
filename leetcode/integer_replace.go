package leetcode

func integerReplacement(n int) int {

	var replace func(n int) int
	replace = func(n int) int {
		if n == 1 {
			return 0
		}
		if n%2 == 0 {
			return replace(n/2) + 1
		}
		return min(replace(n+1), replace(n-1)) + 1
	}
	i := replace(n)
	return i

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

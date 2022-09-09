package leetcode

func minOperations(logs []string) int {
	var ans int
	for _, log := range logs {
		switch log {
		case "../":
			if ans > 0 {
				ans--
			}
		case "./":
		default:
			ans++
		}
	}
	return ans
}

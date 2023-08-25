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

func minOperationsStr(s string) int {
	odd := []string{}
	even := []string{}
	for i, ele := range s {
		if i%2 == 0 {
			even = append(even, string(ele))
			continue
		}
		odd = append(odd, string(ele))
	}
	i := opt(odd, "1") + opt(even, "0")
	j := opt(odd, "0") + opt(even, "1")
	return min(i, j)
}

func opt(s []string, target string) int {
	var ans int
	for _, ele := range s {
		if ele != target {
			ans++
		}
	}
	return ans
}

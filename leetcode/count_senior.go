package leetcode

import "strconv"

func countSeniors(details []string) int {
	var count int
	for _, e := range details {
		ageStr := e[11:13]
		age, _ := strconv.Atoi(ageStr)
		if age > 60 {
			count++
		}

	}
	return count

}

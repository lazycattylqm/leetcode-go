package leetcode

func countWords(words1 []string, words2 []string) int {
	m1 := make(map[string]int)
	m2 := make(map[string]int)
	for _, w := range words1 {
		m1[w]++
	}
	for _, w := range words2 {
		m2[w]++
	}
	for k, v := range m1 {
		if v > 1 {
			delete(m1, k)
		}
	}

	for k, v := range m2 {
		if v > 1 {
			delete(m2, k)
		}
	}
	res := 0
	for w, c := range m1 {
		if c == 1 && m2[w] == 1 {
			res++
		}
	}

	return res

}

package leetcode

func checkIfPangram(sentence string) bool {
	makeMap := make(map[string]bool)
	for _, e := range sentence {
		makeMap[string(e)] = true

	}
	return len(makeMap) == 26
}

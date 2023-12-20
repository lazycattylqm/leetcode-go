package leetcode

func isAcronym(words []string, s string) bool {
	join:=""
	for _, word := range words {
		join += string(word[0])
	}
	return join == s
}
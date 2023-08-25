package leetcode

import "regexp"

func numJewelsInStones(jewels string, stones string) int {
	regexString := "[" + jewels + "]"
	compile := regexp.MustCompile(regexString)
	allString := compile.ReplaceAllString(stones, "")
	return len(stones) - len(allString)

}

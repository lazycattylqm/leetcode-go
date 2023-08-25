package leetcode

func checkString(s string) bool {
	res := 0
	mark := "a"
	for _, i := range s {
		if string(i) != mark {
			res++
			mark = string(i)
		}
		if res >= 2 {
			return false
		}
	}
	return true

}

func CheckString(s string) bool {
	return checkString(s)
}

package leetcode

func countVowelStrings(n int) int {
	array := []int{1, 1, 1, 1, 1}

	if n == 1 {
		sum := 0
		for _, e := range array {
			sum += e
		}
		return sum
	}

	for i := 1; i < n; i++ {
		tempArray := make([]int, 5)
		for j := 0; j < 5; j++ {
			sum := 0
			for k := 0; k <= j; k++ {
				sum += array[k]
			}
			tempArray[j] = sum
		}
		array = tempArray
	}

	sum := 0
	for _, e := range array {
		sum += e
	}

	return sum

}

package leetcode

func MaxAbsoluteSum(nums []int) int {
	maxSum := getMaxSum(nums)
	minSum := getMinSum(nums)
	return max(abs(maxSum), abs(minSum))
}

func getMaxSum(nums []int) int {
	result := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		number := nums[i]
		sum = max(number, sum+number)
		result = max(result, sum)
	}
	return result
}

func getMinSum(nums []int) int {
	result := nums[0]
	sum := nums[0]
	for i := 1; i < len(nums); i++ {
		number := nums[i]
		sum = min(number, sum+number)
		result = min(result, sum)
	}
	return result
}

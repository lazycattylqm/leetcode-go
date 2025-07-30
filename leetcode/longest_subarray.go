package leetcode

func longestSubarray(nums []int) int {
	maxValue := nums[0]
	ans, cnt := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxValue {
			ans, cnt = 1, 1
			maxValue = nums[i]
		} else if nums[i] < maxValue {
			cnt = 0
		} else if nums[i] == maxValue {
			cnt++
		}
		ans = max(ans, cnt)
	}
	return ans
}

package leetcode

import "sort"

func countPairs(nums []int, target int) int {
	sort.Ints(nums)
	left := 0
	right := len(nums) - 1
	count := 0
	for left < right {
		if nums[left]+nums[right] >= target {
			right--
		} else {
			count += (right - left)
			left++
		}

	}
	return count

}

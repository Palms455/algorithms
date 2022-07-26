package main

func search(nums []int, target int) int {
	m, left, right := 0, 0, len(nums)-1

	for left <= right {
		m = (left + right) / 2
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			right = m - 1
			continue
		}
		left = m + 1
	}
	return -1
}

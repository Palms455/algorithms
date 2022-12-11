package main

import "math/rand"

func partition(nums []int, i, j int) int {
	pivot := rand.Intn(j-i) + i
	nums[pivot], nums[i] = nums[i], nums[pivot]
	low_idx := i
	pivot = i
	for k := i + 1; k < j; k++ {
		if nums[k] < nums[pivot] {
			low_idx++
			if low_idx != k {
				nums[low_idx], nums[k] = nums[k], nums[low_idx]
			}
		}
	}
	nums[pivot], nums[low_idx] = nums[low_idx], nums[pivot]
	return low_idx
}

func quickSort(nums []int, i, j int) {
	if i == j {
		return
	}
	pivot := partition(nums, i, j)
	quickSort(nums, i, pivot)
	quickSort(nums, pivot+1, j)
}

func maximumGap(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	quickSort(nums, 0, len(nums))
	cnt := 0
	for i := 1; i < len(nums); i++ {
		curr := nums[i] - nums[i-1]
		if cnt < curr {
			cnt = curr
		}
	}
	return cnt
}

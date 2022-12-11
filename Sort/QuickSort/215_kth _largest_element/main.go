package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	var quickSort func([]int, int, int)
	var partition func([]int, int, int) int

	partition = func(nums []int, left int, right int) int {
		pivot := rand.Intn(right-left) + left
		nums[left], nums[pivot] = nums[pivot], nums[left]
		low_index := left
		pivot_point := left
		for i := left + 1; i < right; i++ {
			if nums[i] < nums[pivot_point] {
				low_index++
				if low_index != i {
					nums[low_index], nums[i] = nums[i], nums[low_index]
				}
			}
		}
		nums[low_index], nums[pivot_point] = nums[pivot_point], nums[low_index]
		return low_index
	}

	quickSort = func(nums []int, left, right int) {
		if left == right {
			return
		}
		pivot := partition(nums, left, right)
		quickSort(nums, left, pivot)
		quickSort(nums, pivot+1, right)
	}

	quickSort(nums, 0, len(nums))
	fmt.Print(nums)

	return nums[len(nums)-k]
}

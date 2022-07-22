package main

func sortColors(nums []int) {
	for i := 0; i < len(nums); i++ {
		min_idx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[min_idx] > nums[j] {
				nums[min_idx], nums[j] = nums[j], nums[min_idx]
			}
		}
	}

}

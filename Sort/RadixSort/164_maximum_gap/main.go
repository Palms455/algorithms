package main

import "fmt"

func maximumGap(nums []int) int {
	powerOfTen := 1
	arr := make([][]int, 10, 10)
	for pow := 0; pow < 9; pow++ {
		for _, elem := range nums {
			idx := elem / powerOfTen % 10
			arr[idx] = append(arr[idx], elem)
		}
		nums = nums[:0]
		for idx, elem := range arr {
			nums = append(nums, elem...)
			arr[idx] = arr[idx][:0]
		}
		powerOfTen *= 10
	}
	cnt := 0
	fmt.Println(nums)
	for i:= 1; i< len(nums); i++{
		curr := nums[i] - nums[i-1]
		if cnt < curr {
			cnt = curr
		}
	}
	return cnt
}




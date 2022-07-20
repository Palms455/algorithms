package main

func singleNumber(nums []int) int {
	n := 0
	for _, item := range nums {
		n ^= item
	}
	return n
}

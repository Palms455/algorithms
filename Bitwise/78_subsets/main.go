package main

func subsets(nums []int) [][]int {
	resultNums := make([][]int, 0, 1<<len(nums))
	for mask := 0; mask < (1 << len(nums)); mask++ {
		innerNums := []int{}
		for idx, num := range nums {
			if (mask & (1 << idx)) > 0 {
				innerNums = append(innerNums, num)
			}
		}
		resultNums = append(resultNums, innerNums)
	}
	return resultNums
}

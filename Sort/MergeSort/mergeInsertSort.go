package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func insertSort(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		j := i - 1

		for j >= 0 && nums[j] > cur {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = cur
	}
	return nums
}

func mergeInsertSort(nums []int) []int {
	if len(nums) <= 20 {
		return insertSort(nums)
	}
	left := mergeInsertSort(nums[:len(nums)/2])
	right := mergeInsertSort(nums[len(nums)/2:])
	return merge(left, right)
}

func main() {
	lst := make([]int, 100)
	for idx, _ := range lst {
		lst[idx] = rand.Intn(100)
	}
	lst2 := make([]int, 100)
	_ = copy(lst2, lst)
	sort.Ints(lst)
	fmt.Println(reflect.DeepEqual(lst, lst2))
	fmt.Println(reflect.DeepEqual(lst, mergeInsertSort(lst2)))
}

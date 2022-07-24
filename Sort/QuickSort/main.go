package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func partition(nums []int, i, j int) int {
	pivot := rand.Intn(j - i) + i
	nums[pivot], nums[i] = nums[i], nums[pivot]
	low_index := i
	pivot = i
	for k := i + 1; k < j; k++ {
		if nums[k] < nums[pivot] {
			low_index++
			if k != low_index {
				nums[low_index], nums[k] = nums[k], nums[low_index]
			}
		}
	}
	nums[low_index], nums[pivot] = nums[pivot], nums[low_index]
	return low_index
}

func quickSort(nums []int, i, j int) {
	if i == j {
		return
	}
	pivot_point := partition(nums, i, j)
	quickSort(nums, i, pivot_point)
	quickSort(nums, pivot_point+1, j)
}

func customSort(nums []int) {
	quickSort(nums, 0, len(nums))
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
	customSort(lst2)
	fmt.Println(reflect.DeepEqual(lst, lst2))
}

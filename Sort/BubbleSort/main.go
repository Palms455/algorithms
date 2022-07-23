package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		is_sorted := true
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				is_sorted = false
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
		if is_sorted {
			return
		}
	}
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
	bubbleSort(lst2)
	fmt.Println(reflect.DeepEqual(lst, lst2))
}

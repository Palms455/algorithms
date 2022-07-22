package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func selectSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		min_idx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[min_idx] > nums[j] {
				nums[j], nums[min_idx] = nums[min_idx], nums[j]
			}
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
	selectSort(lst2)
	fmt.Println(reflect.DeepEqual(lst, lst2))

}

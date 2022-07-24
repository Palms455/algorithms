package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func radixSort(nums []int) []int {
	power_of_ten := 1
	arr := make([][]int, 10, 10)
	for pow := 0; pow < 5; pow++ {
		for _, elem := range nums {
			idx := elem / power_of_ten % 10
			arr[idx] = append(arr[idx], elem)
		}
		nums = nums[:0]
		for idx, elem := range arr {
			nums = append(nums, elem...)
			arr[idx] = arr[idx][:0]
		}
		power_of_ten *= 10
	}
	return nums
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

	fmt.Println(reflect.DeepEqual(lst, radixSort(lst2)))
}

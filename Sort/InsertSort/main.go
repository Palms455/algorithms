package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func insertSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		cur := nums[i]
		j := i - 1
		for j >= 0 && nums[j] > cur {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = cur
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
	insertSort(lst2)
	sort.Ints(lst)
	fmt.Println(reflect.DeepEqual(lst, lst2))
}

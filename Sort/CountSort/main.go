package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func countingSort(nums []int) {
	var minimum, maximum int
	for _, num := range nums {
		if minimum > num {
			minimum = num
		}
		if maximum < num {
			maximum = num
		}
	}
	cnt := make([]int, maximum-minimum+1)
	for _, num := range nums {
		cnt[num-minimum]++
	}
	nums = nums[:0]
	for idx, count_num := range cnt {
		for i := 0; i < count_num; i++ {
			nums = append(nums, idx+minimum)
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
	countingSort(lst2)
	fmt.Println(reflect.DeepEqual(lst, lst2))
}

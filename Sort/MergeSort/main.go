package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
)

func merge(left []int, right []int) []int {
	var left_idx, right_idx int
	newArr := make([]int, 0, len(left)+len(right))
	for left_idx < len(left) && right_idx < len(right) {
		if left[left_idx] < right[right_idx] {
			newArr = append(newArr, left[left_idx])
			left_idx++
		} else {
			newArr = append(newArr, right[right_idx])
			right_idx++
		}
	}
	newArr = append(newArr, left[left_idx:]...)
	newArr = append(newArr, right[right_idx:]...)
	return newArr
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left := mergeSort(nums[:len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])
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
	fmt.Println(reflect.DeepEqual(lst, mergeSort(lst2)))
}

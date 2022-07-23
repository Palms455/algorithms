package main

func countGoodSubstrings(s string) int {
	var left, nums int
	state := map[int32]int{}
	for idx, elem := range s {
		state[elem]++
		if idx-left == 3 {
			left_elem := int32(s[left])
			state[left_elem]--
			if state[left_elem] == 0 {
				delete(state, left_elem)
			}
			left++
		}
		if len(state) == 3 {
			nums++
		}
	}
	return nums
}

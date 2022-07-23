package main

func lengthOfLongestSubstring(s string) int {
	var length, left int
	state := map[int32]bool{}
	for idx, elem := range s {
	inner:
		for {
			ok := state[elem]
			if ok {
				delete(state, int32(s[left]))
				left++
				continue
			}
			break inner
		}
		state[elem] = true
		if length < idx-left+1 {
			length = idx - left + 1
		}
	}
	return length
}

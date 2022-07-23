package main

func lengthOfLongestSubstring(s string) int {
	var length, left int
	state := map[string]bool{}
	for idx, elem := range s {
		inner:
		for {
			ok := state[string(elem)]
			if ok{
				delete(state, string(s[left]))
				left++
				continue
			}
			break inner
		}
		state[string(elem)] = true
		if length < idx-left+1 {
			length = idx - left + 1
		}
	}
	return length
}
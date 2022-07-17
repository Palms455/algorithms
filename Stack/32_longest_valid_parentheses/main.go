package main

import (
	"fmt"
)

// Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Valid Parentheses.

type Stack []int

func (s *Stack) Top() int {
	if len(*s) == 0 {
		return 0
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() int {
	if len(*s) == 0 {
		return 0
	}
	val := s.Top()
	*s = (*s)[:len(*s)-1]
	return val
}

func longestValidParentheses(s string) int {
	stack := Stack{-1}
	length := 0
	for idx, symbol := range s {
		if string(symbol) == "(" {
			stack = append(stack, idx)
			continue
		}
		if len(stack) == 1 {
			stack[0] = idx
			continue
		}
		stack.Pop()
		val := idx - stack.Top()
		if length < val {
			length = val
		}
	}
	return length
}

func main() {
	fmt.Println(longestValidParentheses("()") == 2)
	fmt.Println(longestValidParentheses(")()(())") == 6)
}

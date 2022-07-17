package main

import "fmt"

//Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.

type Stack []string

func (s *Stack) isEmpty() bool {
	if len(*s) == 0 {
		return true
	}
	return false
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Top() string{
	if s.isEmpty() {
		return ""
	}
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() string {
	if s.isEmpty() {
		return ""
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val
}


func isOpen(symbol string) bool {
	return symbol == "[" || symbol == "{" || symbol == "("
}

func isPair(start_s, end_s string) bool {
	return start_s == "(" && end_s == ")" || start_s == "{" && end_s == "}" || start_s == "[" && end_s == "]"
}

func isValid(s string) bool {
	stack := Stack{}
	for _, symbol := range s {
		if isOpen(string(symbol)) {
			stack.Push(string(symbol))
			continue
		}
		if stack.isEmpty() || !isPair(stack.Pop(), string(symbol)){
			return false
		}
	}
	if stack.isEmpty() {
		return true
	}
	return false
}


func main() {
	example1 := "(){}[]"
	example2 := "({}]"
	fmt.Println(isValid(example1))
	fmt.Println(isValid(example2))
}

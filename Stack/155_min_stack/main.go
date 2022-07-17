package main

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

func (s *Stack) Push(val int) {
	*s = append(*s, val)
}



type MinStack struct {
	BaseStack Stack
	Min Stack
}


func Constructor() MinStack {
	return MinStack{
		BaseStack: Stack{},
		Min: Stack{},
	}
}


func (this *MinStack) Push(val int)  {
	this.BaseStack.Push(val)
	if len(this.Min) == 0 || val <= this.Min.Top() {
		this.Min.Push(val)
	}
}


func (this *MinStack) Pop()  {
	if this.BaseStack.Pop() == this.Min.Top() {
		this.Min.Pop()
	}
}


func (this *MinStack) Top() int {
	return this.BaseStack.Top()
}


func (this *MinStack) GetMin() int {
	return this.Min.Top()
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
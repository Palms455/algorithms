package main

type MyStack struct {
	Stack []int
}

func Constructor() MyStack {
	return MyStack{
		Stack: []int{},
	}
}

func (this *MyStack) Push(x int) {
	this.Stack = append(this.Stack, x)
}

func (this *MyStack) Top() int {
	if this.Empty() {
		return 0
	}
	return this.Stack[len(this.Stack)-1]
}

func (this *MyStack) Pop() int {
	if this.Empty() {
		return 0
	}
	val := this.Top()
	this.Stack = this.Stack[:len(this.Stack)-1]
	return val
}

func (this *MyStack) Empty() bool {
	return len(this.Stack) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

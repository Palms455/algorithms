package main

type MyQueue struct {
	Queue []int
}

func Constructor() MyQueue {
	return MyQueue{
		Queue: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.Queue = append(this.Queue, x)
}

func (this *MyQueue) Pop() int {
	if !this.Empty() {
		val := this.Queue[0]
		this.Queue = this.Queue[1:]
		return val
	}
	return 0
}

func (this *MyQueue) Peek() int {
	if this.Empty() {
		return 0
	}
	return this.Queue[0]

}

func (this *MyQueue) Empty() bool {
	return len(this.Queue) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

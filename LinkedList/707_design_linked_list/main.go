package main

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{
		head: &Node{
			val: 0,
		},
		size: 0,
	}

}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	prev := this.head
	for i := 0; i <= index; i++ {
		prev = prev.next
	}
	return prev.val
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	prev := this.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	node := &Node{
		val:  val,
		next: prev.next,
	}
	prev.next = node
	this.size++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	prev := this.head
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	this.size--
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */

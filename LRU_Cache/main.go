package main

type Node struct {
	Key   int
	Value int
	Prev  *Node
	Next  *Node
}

type LRUCache struct {
	Capacity int
	tail     *Node
	head     *Node
	memo     map[int]*Node
}

func Constructor(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.Next = tail
	tail.Prev = head
	return LRUCache{
		Capacity: capacity,
		tail:     tail,
		head:     head,
		memo:     make(map[int]*Node, capacity),
	}
}

func (this *LRUCache) append(node *Node) {
	prev := this.tail.Prev
	prev.Next = node
	node.Prev = prev
	node.Next = this.tail
	this.tail.Prev = node
}

func (this *LRUCache) remove(node *Node) {
	prev := node.Prev
	next := node.Next
	prev.Next = next
	next.Prev = prev
}

func (this *LRUCache) removeFirst() {
	node := this.head.Next
	this.remove(node)
	delete(this.memo, node.Key)
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.memo[key]
	if !ok {
		return -1
	}
	if node.Next != this.tail {
		this.remove(node)
		this.append(node)
	}
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.memo[key]
	if ok {
		node.Value = value
		this.Get(key)
		return
	}
	node = &Node{Value: value, Key: key}
	this.append(node)
	this.memo[key] = node
	if len(this.memo) > this.Capacity {
		this.removeFirst()
	}
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

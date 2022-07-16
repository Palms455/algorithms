package main

type Node struct {
	Val  string
	Next *Node
	Prev *Node
}

type BrowserHistory struct {
	Node *Node
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		Node: &Node{
			Val: homepage,
		},
	}
}

func (this *BrowserHistory) Visit(url string) {
	node := &Node{
		Val:  url,
		Prev: this.Node,
	}
	this.Node.Next = node
	this.Node = node
}

func (this *BrowserHistory) Back(steps int) string {
	for ; steps > 0 && this.Node.Prev != nil; steps-- {
		this.Node = this.Node.Prev
	}
	return this.Node.Val
}

func (this *BrowserHistory) Forward(steps int) string {
	for ; steps > 0 && this.Node.Next != nil; steps-- {
		this.Node = this.Node.Next
	}
	return this.Node.Val
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */

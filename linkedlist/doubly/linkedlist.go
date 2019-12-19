package doubly

type List struct {
	head  *Node
	tail  *Node
	count int
}

type Node struct {
	value int
	prev  *Node
	next  *Node
}

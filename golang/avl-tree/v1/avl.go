package avl

// node
type node struct {
	val         int
	height      int
	left, right *node
}

// insert inserts new node with val to tree
func insert(p *node, val int) *node {
	if p == nil {
		return newNode(val)
	}
	if p.val > val {
		p.left = insert(p.left, val)
	} else if p.val < val {
		p.right = insert(p.right, val)
	}
	return balance(p)
}

// balance
func balance(p *node) *node {
	if balanceFactor(p) == -2 {
		if balanceFactor(p.left) > 0 {
			p.left = leftRotate(p.left)
		}
		return rightRotate(p)
	}
	if balanceFactor(p) == 2 {
		if balanceFactor(p.right) < 0 {
			p.right = rightRotate(p.right)
		}
		return leftRotate(p)
	}
	return p
}

// rightRotate
func rightRotate(p *node) *node {
	q := p.left
	p.left = q.right
	q.right = p

	fixHeight(p)
	fixHeight(q)

	return q
}

// leftRotate
func leftRotate(p *node) *node {
	q := p.right
	p.right = q.left
	q.left = p

	fixHeight(p)
	fixHeight(q)

	return q
}

// newNode builds new *node with height=1
func newNode(val int) *node {
	return &node{
		val:    val,
		height: 1,
	}
}

// balanceFactor calcs balance factor for p
func balanceFactor(p *node) int {
	var hl, hr int
	if p.left != nil {
		hl = p.left.height
	}
	if p.right != nil {
		hr = p.right.height
	}
	return hr - hl
}

// fixHeight set new height for p
func fixHeight(p *node) {
	var hl, hr int
	if p.left != nil {
		hl = p.left.height
	}
	if p.right != nil {
		hr = p.right.height
	}
	p.height = max(hl, hr) + 1
}

// max
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

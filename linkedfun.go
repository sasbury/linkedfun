package linkedfun

/*
Node represents a simple linked list node with an integer value.
*/
type Node struct {
	value int
	next  *Node
}

//Move 1 step
func slowStep(n *Node) *Node {
	return n.next
}

//Move 2 steps
func fastStep(n *Node) *Node {
	n = n.next

	if n != nil {
		n = n.next
	}

	return n
}

func FindCycle(head *Node) bool {

	//Empty list is never a cycle
	if head == nil {
		return false
	}

	//Start walking right away, if we get nil,nil then it is a short
	//non-cycle list
	slowWalker := slowStep(head)
	fastWalker := fastStep(head)

	//If walkers are nil they hit the end, so no cycle
	for slowWalker != nil && fastWalker != nil {

		//If walkers are equal they hit a cycle
		if slowWalker == fastWalker {
			return true
		}

		//Move the walkers forward
		slowWalker = slowStep(slowWalker)
		fastWalker = fastStep(fastWalker)
	}

	return false
}

func Reverse(head *Node) *Node {

	h := head
	n := head.next
	h.next = nil

	for n != nil {
		newN := n.next
		n.next = h
		h = n
		n = newN
	}

	return h
}

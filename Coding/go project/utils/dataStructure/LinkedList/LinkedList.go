package models

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (linkedList LinkedList) Iterate(visitor func(current *Node, next *Node, previous *Node)) {
	current := linkedList.Head
	if current == nil {
		return
	}
	var previous *Node
	for {
		visitor(current, current.Next, previous)
		if current.Next == nil {
			return
		}
		previous = current
		current = current.Next
	}
}

func FromVariadic(values ...interface{}) *LinkedList {
	linkedList := &LinkedList{}
	var previous *Node
	for _, value := range values {
		node := &Node{Value: value}
		if linkedList.Head == nil {
			linkedList.Head = node
			previous = node
			continue
		}
		linkedList.Tail = node
		if previous != nil {
			previous.Next = node
		}
		previous = node
	}
	return linkedList
}

package list

/*
In computer science, a linked list is a linear collection of data elements,
in which linear order is not given by their physical placement in memory.
Each pointing to the next node by means of a pointer. It is a data structure
consisting of a group of nodes which together represent a sequence.
Here is source code of the Go Program to Implement Single Unsorted Linked List
*/

type node struct {
	key  interface{}
	prev *node
	next *node
}

type LinkedList struct {
	head *node
}

func (l *LinkedList) Insert(key interface{}) {
	newHead := &node{
		key:  key,
		next: l.head,
	}

	if l.head != nil {
		l.head.prev = newHead
	}
	l.head = newHead
}

func (l *LinkedList) Reverse() {
	if l.head == nil {
		return
	}

	var temp *node
	cursor := l.head
	for cursor != nil {
		newPrev := cursor.next
		cursor.next = temp
		temp = cursor
		cursor = newPrev
	}

	l.head = temp
}

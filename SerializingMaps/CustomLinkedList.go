package main

import "fmt"

// Define Node structure
type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

// Define LinkedList structure
type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

// Create a new linked list with a nil head and tail
func CreateLinkedList[T any]() *LinkedList[T] {
	linkedList := new(LinkedList[T]) // Using new() for heap allocation
	linkedList.head = new(Node[T])   // Sentinel head node
	linkedList.tail = new(Node[T])   // Sentinel tail node

	linkedList.head.next = linkedList.tail
	linkedList.tail.prev = linkedList.head

	return linkedList
}

// Create a new node
func CreateNode[T any](value T, next *Node[T], prev *Node[T]) *Node[T] {
	return &Node[T]{
		value: value,
		next:  next,
		prev:  prev,
	}
}

// Append value at the head of the list
func (linkedList *LinkedList[T]) AppendHead(value T) {
	newNode := CreateNode(value, linkedList.head.next, linkedList.head)
	linkedList.head.next.prev = newNode
	linkedList.head.next = newNode
}

// Append value at the tail of the list
func (linkedList *LinkedList[T]) AppendTail(value T) {
	newNode := CreateNode(value, linkedList.tail, linkedList.tail.prev)
	linkedList.tail.prev.next = newNode
	linkedList.tail.prev = newNode
}

// Check if the linked list is empty
func (linkedList *LinkedList[T]) IsEmpty() bool {
	return linkedList.head.next == linkedList.tail
}

// Delete a node from the linked list
func (linkedList *LinkedList[T]) Delete(node *Node[T]) {
	if node == nil || node == linkedList.head || node == linkedList.tail {
		return // Don't delete head/tail sentinels or nil node
	}
	node.prev.next = node.next
	node.next.prev = node.prev
}

// Remove and return the value at the head
func (linkedList *LinkedList[T]) PopHead() (T, bool) {
	if linkedList.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	node := linkedList.head.next
	linkedList.Delete(node)
	return node.value, true
}

// Remove and return the value at the tail
func (linkedList *LinkedList[T]) PopTail() (T, bool) {
	if linkedList.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	node := linkedList.tail.prev
	linkedList.Delete(node)
	return node.value, true
}

// Peek at the value at the head without removing it
func (linkedList *LinkedList[T]) PeekHead() (T, bool) {
	if linkedList.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return linkedList.head.next.value, true
}

// Peek at the value at the tail without removing it
func (linkedList *LinkedList[T]) PeekTail() (T, bool) {
	if linkedList.IsEmpty() {
		var zeroValue T
		return zeroValue, false
	}
	return linkedList.tail.prev.value, true
}

// Print the linked list for debugging
func (linkedList *LinkedList[T]) Print() {
	current := linkedList.head.next
	for current != linkedList.tail {
		fmt.Print(current.value, " <-> ")
		current = current.next
	}
	fmt.Println("nil")
}

// Main function for testing
func main() {
	ll := CreateLinkedList[int]()
	ll.AppendHead(10)
	ll.AppendHead(20)
	ll.AppendTail(30)
	ll.AppendTail(40)

	ll.Print() // Output: 20 <-> 10 <-> 30 <-> 40 <-> nil

	value, found := ll.PeekHead()
	fmt.Println("Peek Head:", value, "Found:", found)
	value, found = ll.PeekTail()
	fmt.Println("Peek Tail:", value, "Found:", found)

	ll.PopHead()
	ll.Print() // Output: 10 <-> 30 <-> 40 <-> nil

	ll.PopTail()
	ll.Print() // Output: 10 <-> 30 <-> nil

	fmt.Println("Is Empty:", ll.IsEmpty()) // Output: false

	ll.PopHead()
	ll.PopHead()
	fmt.Println("Is Empty:", ll.IsEmpty()) // Output: true
}

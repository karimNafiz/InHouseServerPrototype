package in_house_linked_list

import "fmt"

// Define Node structure
type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

// Define LinkedList structure
type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	len  int
}

// Create a new linked list with a nil Head and Tail
func CreateLinkedList[T any]() *LinkedList[T] {
	linkedList := new(LinkedList[T]) // Using new() for heap allocation
	linkedList.Head = new(Node[T])   // Sentinel Head node
	linkedList.Tail = new(Node[T])   // Sentinel Tail node

	linkedList.Head.Next = linkedList.Tail
	linkedList.Tail.Prev = linkedList.Head

	return linkedList
}

// Create a new node
func CreateNode[T any](Value T, Next *Node[T], Prev *Node[T]) *Node[T] {
	return &Node[T]{
		Value: Value,
		Next:  Next,
		Prev:  Prev,
	}
}

// Append Value at the Head of the list
func (linkedList *LinkedList[T]) AppendHead(Value T) {
	newNode := CreateNode(Value, linkedList.Head.Next, linkedList.Head)
	linkedList.Head.Next.Prev = newNode
	linkedList.Head.Next = newNode
	linkedList.len++
}

// Append Value at the Tail of the list
func (linkedList *LinkedList[T]) AppendTail(Value T) {
	newNode := CreateNode(Value, linkedList.Tail, linkedList.Tail.Prev)
	linkedList.Tail.Prev.Next = newNode
	linkedList.Tail.Prev = newNode
	linkedList.len++
}

// Check if the linked list is empty
func (linkedList *LinkedList[T]) IsEmpty() bool {
	return linkedList.Head.Next == linkedList.Tail
}

// Delete a node from the linked list
func (linkedList *LinkedList[T]) Delete(node *Node[T]) {
	if node == nil || node == linkedList.Head || node == linkedList.Tail {
		return // Don't delete Head/Tail sentinels or nil node
	}
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
	linkedList.len--
}

// Remove and return the Value at the Head
func (linkedList *LinkedList[T]) PopHead() (T, bool) {
	if linkedList.IsEmpty() {
		var zero_value T
		return zero_value, false
	}
	node := linkedList.Head.Next
	linkedList.Delete(node)
	return node.Value, true
}

// Remove and return the Value at the Tail
func (linkedList *LinkedList[T]) PopTail() (T, bool) {
	if linkedList.IsEmpty() {
		var zero_value T
		return zero_value, false
	}
	node := linkedList.Tail.Prev
	linkedList.Delete(node)
	return node.Value, true
}

// Peek at the Value at the Head without removing it
func (linkedList *LinkedList[T]) PeekHead() (T, bool) {
	if linkedList.IsEmpty() {
		var zero_value T
		return zero_value, false
	}
	return linkedList.Head.Next.Value, true
}

// Peek at the Value at the Tail without removing it
func (linkedList *LinkedList[T]) PeekTail() (T, bool) {
	if linkedList.IsEmpty() {
		var zero_value T
		return zero_value, false
	}
	return linkedList.Tail.Prev.Value, true
}

func (linkedList *LinkedList[T]) GetLength() int {
	return linkedList.len
}

func (linkedList *LinkedList[T]) GetIndex(index int) (T, bool) {
	var zero_value T
	if linkedList.IsEmpty() || linkedList.GetLength() <= index {
		return zero_value, false
	}
	current := linkedList.Head.Next
	element_index := 0
	for current != linkedList.Tail {
		if index == element_index {
			return current.Value, true
		}
		element_index++
		current = current.Next
	}
	return zero_value, false

}

// Print the linked list for debugging
func (linkedList *LinkedList[T]) Print() {
	current := linkedList.Head.Next
	for current != linkedList.Tail {
		fmt.Print(current.Value, " <-> ")
		current = current.Next
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

	Value, found := ll.PeekHead()
	fmt.Println("Peek Head:", Value, "Found:", found)
	Value, found = ll.PeekTail()
	fmt.Println("Peek Tail:", Value, "Found:", found)

	ll.PopHead()
	ll.Print() // Output: 10 <-> 30 <-> 40 <-> nil

	ll.PopTail()
	ll.Print() // Output: 10 <-> 30 <-> nil

	fmt.Println("Is Empty:", ll.IsEmpty()) // Output: false

	ll.PopHead()
	ll.PopHead()
	fmt.Println("Is Empty:", ll.IsEmpty()) // Output: true
}

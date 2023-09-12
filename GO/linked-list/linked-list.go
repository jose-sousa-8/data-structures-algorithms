package main

import (
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func createNode[T any](value T) *Node[T] {
	return &Node[T]{value: value, next: nil}
}

func createLinkedList[T comparable](value T) *LinkedList[T] {
	node := createNode[T](value)
	return &LinkedList[T]{head: node, tail: node, length: 1}
}

func (ll *LinkedList[T]) append(value T) {
	node := createNode[T](value)
	if ll.length == 0 {
		ll.head = node
		ll.tail = node
	} else {
		ll.tail.next = node
		ll.tail = node
	}

	ll.length++
}

func (ll *LinkedList[T]) prepend(value T) {
	node := createNode[T](value)
	if ll.length == 0 {
		ll.head = node
		ll.tail = node
	} else {
		node.next = ll.head
		ll.head = node
	}
	ll.length++
}

func (ll *LinkedList[T]) get(index int) *Node[T] {
	if index < 0 || index >= ll.length {
		return nil
	}

	counter := 0
	tmp := ll.head
	for counter < index {
		tmp = tmp.next
		counter++
	}

	return tmp
}

func (ll *LinkedList[T]) set(value T, index int) bool {
	node := ll.get(index)
	node.value = value
	return true
}

func (ll *LinkedList[T]) insert(value T, index int) bool {
	if index < 0 || index > ll.length {
		return false
	}

	if index == 0 {
		ll.prepend(value)
	} else if index == ll.length-1 {
		ll.append(value)
	} else {
		newNode := createNode[T](value)
		tmp := ll.get(index - 1)
		newNode.next = tmp.next
		tmp.next = newNode
		ll.length++
	}

	return true
}

func (ll *LinkedList[T]) deleteLast() {
	if ll.length == 0 {
		return
	}
	if ll.length == 1 {
		ll.head = nil
		ll.tail = nil
	} else {
		tmp := ll.head
		previous := tmp
		for tmp.next != nil {
			previous = tmp
			tmp = tmp.next
		}
		ll.tail = previous
		ll.tail.next = nil
	}
	ll.length--
}

func (ll *LinkedList[T]) deleteFirst() {
	if ll.length == 0 {
		return
	}
	if ll.length == 1 {
		ll.head = nil
		ll.tail = nil
	} else {
		ll.head = ll.head.next
	}
	ll.length--
}

func printLinkedList[T comparable](linkedList *LinkedList[T]) {
	node := linkedList.head
	for node != nil {
		fmt.Println(node.value)
		node = node.next
	}
}

func (ll *LinkedList[T]) deleteNode(index int) {
	if index < 0 || index >= ll.length {
		return
	}

	if index == 0 {
		ll.deleteFirst()
	} else if index == ll.length-1 {
		ll.deleteLast()
	} else {
		previous := ll.get(index - 1)
		tmp := previous.next
		previous.next = tmp.next
		ll.length--
	}
}

func (ll *LinkedList[T]) reverse() {
	tmp := ll.head
	ll.head = ll.tail
	ll.tail = tmp

	var before *Node[T]
	for i := 0; i < ll.length; i++ {
		after := tmp.next
		tmp.next = before
		before = tmp
		tmp = after
	}
}

// Initialize two pointers, slow and fast, both pointing to the head of the linked list.
// Traverse the linked list using a while loop.
// The loop continues as long as fast is not nullptr (i.e., it has not reached the end of the list),
// and fast->next is also not nullptr (i.e., there is at least one more node after the current fast node).
// Inside the loop, move the slow pointer one step forward (i.e., slow = slow->next)
// and the fast pointer two steps forward (i.e., fast = fast->next->next).
// Since the fast pointer moves twice as fast as the slow pointer,
// by the time the fast pointer reaches the end of the list or goes beyond it, the slow pointer will be at the middle node.
// When the loop terminates, return the slow pointer, which now points to the middle node.
// In the case of an even number of nodes, the fast pointer will reach the end of the list,
// while the slow pointer will point to the first middle node (the one closer to the head).
// For an odd number of nodes, the fast pointer will go beyond the end of the list, and the slow pointer will point to the exact middle node.
func (ll *LinkedList[T]) findMiddleNode() *Node[T] {
	fast := ll.head
	slow := ll.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	// When the loop ends, slow points to the middle
	return slow
}

// Initialize the slow and fast pointers to the head of the linked list.
// Start a while loop that continues as long as the fast pointer is not null and the next node of the fast pointer is not null.
// The loop will break if either of these conditions is false, indicating that the end of the list has been reached,
// and thus there is no loop.
// Inside the loop, move the slow pointer one step forward (to the next node)
// and the fast pointer two steps forward (to the next-to-next node).
// Check if the slow pointer is equal to the fast pointer.
// If they are equal, it means the pointers have met inside the loop, so a loop exists in the linked list.
// In this case, return true.
// If the while loop completes without the pointers meeting, it means there is no loop in the list.
// In this case, return false.
// This algorithm has a time complexity of O(n) and a space complexity of O(1),
//  making it an efficient way to detect loops in a linked list.

func (ll *LinkedList[T]) hasLoop() bool {
	if ll.head == nil {
		return false
	}

	slow := ll.head
	fast := ll.head.next
	for slow != fast {
		if fast == nil || fast.next == nil {
			return false
		}
		slow = slow.next
		fast = fast.next.next
	}
	return true
}

// Initialize two pointers, slow and fast, both pointing to the head of the list.
// Move the fast pointer k steps ahead in the list.
// 	For each step from 0 to k-1, perform the following:
// 		If the fast pointer is nullptr, return nullptr (k is larger than the list size).
// 		Move the fast pointer to the next node.

// Traverse the list with both slow and fast pointers until the fast pointer reaches the end.
//
//	While the fast pointer is not nullptr, perform the following:
//		Move the slow pointer to the next node.
//		Move the fast pointer to the next node.
//
// Return the slow pointer, which is now at the kth node from the end of the list.
func (ll *LinkedList[T]) findKthFromEnd(k int) *Node[T] {
	slow := ll.head
	fast := ll.head

	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.next
	}

	for fast != nil {
		slow = slow.next
		fast = fast.next
	}

	return slow
}

func (ll *LinkedList[T]) removeDuplicates() {
	if ll.length == 0 || ll.length == 1 {
		return
	}

	m := map[T]bool{}

	previous := &Node[T]{}
	current := ll.head
	for current != nil {
		if m[current.value] {
			previous.next = current.next
			current = previous.next
			ll.length -= 1
		} else {
			m[current.value] = true
			previous = current
			current = current.next
		}
	}
}

func (ll *LinkedList[T]) reverseBetween(m, n int) {
	if ll.head == nil {
		return
	}

	var d T
	dummy := createNode(d)
	dummy.next = ll.head
	prev := dummy

	for i := 0; i < m; i++ {
		prev = prev.next
	}

	current := prev.next
	for i := 0; i < n-m; i++ {
		temp := current.next
		current.next = temp.next
		temp.next = prev.next
		prev.next = temp
	}

	ll.head = dummy.next
}

func main() {
	var linkedList = createLinkedList[int](1)
	linkedList.append(2)
	linkedList.append(2)
	linkedList.append(3)
	linkedList.append(3)
	linkedList.append(4)
	linkedList.append(5)
	linkedList.removeDuplicates()
	printLinkedList(linkedList)
	linkedList.reverseBetween(0, 1)
	printLinkedList(linkedList)
}

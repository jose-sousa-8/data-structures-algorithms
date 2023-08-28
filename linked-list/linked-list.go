package main

import (
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	length int
}

func createNode[T any](value T) *Node[T] {
	return &Node[T]{value: value, next: nil}
}

func createLinkedList[T any](value T) *LinkedList[T] {
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

func (ll *LinkedList[T]) get(index int) (*Node[T], error) {
	if index < 0 || index >= ll.length {
		return &Node[T]{}, fmt.Errorf("index out of bounnds")
	}

	counter := 0
	tmp := ll.head
	for counter < index {
		tmp = tmp.next
		counter++
	}

	return tmp, nil
}

func (ll *LinkedList[T]) set(value T, index int) bool {
	node, err := ll.get(index)
	if err != nil {
		return false
	}
	node.value = value
	return true
}

func (ll *LinkedList[T]) insert(value T, index int) bool {
	if index < 0 || index >= ll.length {
		return false
	}

	if index == 0 {
		ll.prepend(value)
	} else if index == ll.length {
		ll.append(value)
	} else {
		newNode := createNode[T](value)
		tmp, _ := ll.get(index - 1)
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

func printLinkedList[T any](linkedList *LinkedList[T]) {
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
		previous, _ := ll.get(index - 1)
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

func main() {
	var linkedList = createLinkedList[int](1)
	printLinkedList[int](linkedList)

	fmt.Println("-----------")

	linkedList.append(5)
	linkedList.append(8)
	printLinkedList[int](linkedList)

	fmt.Println("-----------")

	linkedList.deleteLast()
	printLinkedList[int](linkedList)

	fmt.Println("-----------")

	linkedList.prepend(8)
	printLinkedList[int](linkedList)

	fmt.Println("-----------")

	linkedList.deleteFirst()
	printLinkedList[int](linkedList)

	fmt.Println("-----------")

	fmt.Println(linkedList.get(1))
	fmt.Println(linkedList.get(2))
	fmt.Println(linkedList.get(0))

	fmt.Println("-----------")

	linkedList.set(10, 0)
	printLinkedList[int](linkedList)

	fmt.Println("-----------")

	linkedList.insert(7, 1)
	printLinkedList[int](linkedList)

	fmt.Println("-----------")
	linkedList.deleteNode(1)
	printLinkedList[int](linkedList)

	fmt.Println("-----------")
	linkedList.append(1)
	linkedList.append(2)
	linkedList.append(3)
	linkedList.append(4)
	printLinkedList[int](linkedList)
	fmt.Println("-----------")
	linkedList.reverse()
	printLinkedList[int](linkedList)
}

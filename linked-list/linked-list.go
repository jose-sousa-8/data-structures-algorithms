package main

import "fmt"

type linkedList[T any] struct {
	head *T
	val  T
	next *T
}

func main() {
	fmt.Print("Hey linked list")
	var ll linkedList[int]
	var head = 2
	ll.head = &head
	ll.val = head
	ll.next = nil
	fmt.Print(ll.head)
}

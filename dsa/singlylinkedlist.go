package main

import "fmt"

type LinkedListNode struct {
	data int
	next *LinkedListNode
}

type LinkedList struct {
	head *LinkedListNode
}

// Adds a new node to the linked list
func (ll *LinkedList) AddToHead(data int){
	newNode := LinkedListNode{data:data}
	newNode.next = ll.head

	ll.head = &newNode
}

func (ll *LinkedList) AddToEnd(data int){
	newNode := LinkedListNode{data:data}
	current := ll.head
	// If there's no head - the new node becomes the head
	if current == nil {
		ll.head = &newNode
		return
	}
	// Travel to the end and then insert the new node
	for current.next != nil {
		current = current.next
	}
	current.next = &newNode
}
// Prints the items in a linked list
func (ll *LinkedList) PrintList(){
	curr := ll.head

	for curr != nil {
		if curr.next != nil {
			fmt.Printf("%d->", curr.data)
		} else {
			fmt.Printf("%d", curr.data)
		}
		curr = curr.next
	}
	fmt.Println()
}

// Reverses a linked list
func (ll *LinkedList) ReverseList() *LinkedListNode {
	var curr, prev, temp *LinkedListNode
	curr = ll.head
	temp = nil 
	prev = nil 

	for curr != nil {
		temp = curr.next
		curr.next = prev

		prev = curr
		curr = temp
	}
	return prev
}
// func main(){
// 	list := &LinkedList{}
// 	list.AddToHead(10)
// 	list.AddToHead(20)
// 	list.AddToHead(30)

// 	list.PrintList()
// }

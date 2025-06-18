package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 * Val int
 * Next *ListNode
 * }
 */

 func deleteMiddle(head *ListNode) *ListNode {
	 if head == nil || head.Next == nil {
		 return nil
	 }
 
	 size := 0
	 current := head
	 // Get the size of the linked list
	 for current != nil {
		 current = current.Next
		 size++
	 }
	 // Get the middle index 
	 targetMidIndex := size / 2
	 // Use slow and fast pointers to keep track of previous and current nodes 
	 prev := head
	 current = head.Next
	 // Traverse the list from beginning to target index,
	 // Adjusting the nodes to perform the deletions
	 for i := 0; i < targetMidIndex-1; i++ {
		 prev = prev.Next
		 current = current.Next
	 }
	 prev.Next = current.Next
 
	 return head
 }
 
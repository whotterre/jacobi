package main

/*
LeetCode 206 - Reverse Linked List
Given the head of a singly linked list, reverse the list, and return the reversed list.
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var curr, prev, temp *ListNode
	curr = head
	temp = nil
	prev = nil

	for curr != nil {
		temp = curr.Next
		curr.Next = prev

		prev = curr
		curr = temp
	}
	return prev
}

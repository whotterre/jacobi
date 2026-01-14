package main

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

func hasCycle(head *ListNode) bool {
	// floyd's turtle and hare
	if head == nil {
		return false
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		// what was the ending condition again?
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}
	return false
}

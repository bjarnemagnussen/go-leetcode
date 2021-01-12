package problem0024

import (
	. "github.com/bjarnemagnussen/go-leetcode/library"
)

func swapPairs(head *ListNode) *ListNode {
	// Solution modifying the values in the list's nodes
	// for cur := head; cur != nil && cur.Next != nil; cur = cur.Next.Next {
	// 	cur.Val, cur.Next.Val = cur.Next.Val, cur.Val
	// }

	var last *ListNode
	for cur := head; cur != nil && cur.Next != nil; cur = cur.Next {
		if last == nil {
			head = cur.Next
		} else {
			last.Next = cur.Next
		}
		last = cur
		cur.Next, cur.Next.Next = cur.Next.Next, cur
	}

	return head
}

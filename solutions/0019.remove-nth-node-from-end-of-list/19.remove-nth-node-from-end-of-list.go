package problem0019

import (
	"github.com/bjarnemagnussen/go-leetcode/library"
)

// ListNode is a node for singly-linked list.
type ListNode = library.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var parent *ListNode
	for i, cur := n, head; cur.Next != nil; i, cur = i-1, cur.Next {
		switch {

		case i > 1:
			continue

		case i == 1:
			parent = head
			continue

		default:
			parent = parent.Next

		}
	}

	if parent == nil {
		return head.Next
	}

	parent.Next = parent.Next.Next

	return head
}

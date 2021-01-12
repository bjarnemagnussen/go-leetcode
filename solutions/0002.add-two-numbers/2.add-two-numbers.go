package problem0002

import (
	"github.com/bjarnemagnussen/go-leetcode/library"
)

// ListNode is a node for singly-linked list.
type ListNode = library.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res
	var leftover int
	for l1 != nil || l2 != nil || leftover != 0 {
		v := leftover
		if l1 != nil {
			v += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v += l2.Val
			l2 = l2.Next
		}
		leftover = v / 10
		cur.Next = &ListNode{Val: v % 10}
		cur = cur.Next
	}
	return res.Next
}

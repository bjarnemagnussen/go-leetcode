package problem0002

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	list1 *ListNode
	list2 *ListNode
}

func makeListNode(in []int) *ListNode {
	if len(in) == 0 {
		return nil
	}

	res := &ListNode{Val: in[0]}
	cur := res
	for i := 1; i < len(in); i++ {
		cur.Next = &ListNode{Val: in[i]}
		cur = cur.Next
	}
	return res
}

func Test_addTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example",
			args: args{
				list1: makeListNode([]int{2, 4, 3}),
				list2: makeListNode([]int{5, 6, 4}),
			},
			want: makeListNode([]int{7, 0, 8}),
		},
		{
			name: "leftover",
			args: args{
				list1: makeListNode([]int{2, 4, 8}),
				list2: makeListNode([]int{5, 6, 4}),
			},
			want: makeListNode([]int{7, 0, 3, 1}),
		},
		{
			name: "uneven",
			args: args{
				list1: makeListNode([]int{1, 4}),
				list2: makeListNode([]int{5, 6, 4}),
			},
			want: makeListNode([]int{6, 0, 5}),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, addTwoNumbers(tt.args.list1, tt.args.list2))
		})
	}
}

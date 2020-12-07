package problem0019

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	n    int
	head *ListNode
}

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example1",
			args: args{
				n:    2,
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}},
			},
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}},
		},
		{
			name: "example2",
			args: args{
				n:    1,
				head: &ListNode{Val: 1},
			},
			want: nil,
		},
		{
			name: "example3",
			args: args{
				n:    1,
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			},
			want: &ListNode{Val: 1},
		},
		{
			name: "front",
			args: args{
				n:    2,
				head: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			},
			want: &ListNode{Val: 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeNthFromEnd(tt.args.head, tt.args.n))
		})
	}
}

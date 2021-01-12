package problem0024

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// Imports ListNode from library
	. "github.com/bjarnemagnussen/go-leetcode/library"
)

type args struct {
	head *ListNode
}

func Test_swapPairs(t *testing.T) {
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}},
			},
			want: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}},
		},
		{
			name: "example2",
			args: args{
				head: nil,
			},
			want: nil,
		},
		{
			name: "example3",
			args: args{
				head: &ListNode{1, nil},
			},
			want: &ListNode{1, nil},
		},
		{
			name: "own1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			},
			want: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, swapPairs(tt.args.head))
		})
	}
}

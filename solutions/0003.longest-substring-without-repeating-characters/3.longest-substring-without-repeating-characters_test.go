package problem0003

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input string
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				input: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "example 2",
			args: args{
				input: "bbbbb",
			},
			want: 1,
		},
		{
			name: "example 3",
			args: args{
				input: "pwwkew",
			},
			want: 3,
		},
		{
			name: "empty",
			args: args{
				input: "",
			},
			want: 0,
		},
		{
			name: "single",
			args: args{
				input: "a",
			},
			want: 1,
		},
		{
			name: "multi 1",
			args: args{
				input: "aaab",
			},
			want: 2,
		},
		{
			name: "multi 2",
			args: args{
				input: "abbab",
			},
			want: 2,
		},
		{
			name: "multi 3",
			args: args{
				input: "abab",
			},
			want: 2,
		},
		{
			name: "random 1",
			args: args{
				input: "gwrnluyjzizjhntmrqfyvxhmcqrkesaiasbyblyn",
			},
			want: 14,
		},
		{
			name: "random 2",
			args: args{
				input: "sgpqcdunblxhihplcjcadaszyuvowltwjlldvleg",
			},
			want: 13,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, lengthOfLongestSubstring(tt.args.input))
		})
	}
}

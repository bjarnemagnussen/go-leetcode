package problem1370

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input string
}

func Test_sortString(t *testing.T) {
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example1",
			args: args{
				input: "aaaabbbbcccc",
			},
			want: "abccbaabccba",
		},
		{
			name: "example2",
			args: args{
				input: "rat",
			},
			want: "art",
		},
		{
			name: "example3",
			args: args{
				input: "leetcode",
			},
			want: "cdelotee",
		},
		{
			name: "example4",
			args: args{
				input: "ggggggg",
			},
			want: "ggggggg",
		},
		{
			name: "example5",
			args: args{
				input: "spo",
			},
			want: "ops",
		},
		{
			name: "empty string",
			args: args{
				input: "",
			},
			want: "",
		},
		{
			name: "single character",
			args: args{
				input: "a",
			},
			want: "a",
		},
		{
			name: "stays same",
			args: args{
				input: "ab",
			},
			want: "ab",
		},
		{
			name: "reverse",
			args: args{
				input: "ba",
			},
			want: "ab",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, sortString(tt.args.input))
		})
	}
}

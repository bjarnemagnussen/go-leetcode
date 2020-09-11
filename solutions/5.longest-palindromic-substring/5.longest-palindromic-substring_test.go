package problem0005

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input string
}

func Test_longestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "example 1",
			args: args{
				input: "babad",
			},
			want: "bab",
		},
		{
			name: "example 2",
			args: args{
				input: "cbbd",
			},
			want: "bb",
		},
		{
			name: "baabab",
			args: args{
				input: "baabab",
			},
			want: "baab",
		},
		{
			name: "a",
			args: args{
				input: "a",
			},
			want: "a",
		},
		{
			name: "kpmmjpn",
			args: args{
				input: "kpmmjpn",
			},
			want: "mm",
		},
		{
			name: "kbajmamjabcpn",
			args: args{
				input: "kbajmamjabcpn",
			},
			want: "bajmamjab",
		},
		{
			name: "abkzcdc",
			args: args{
				input: "abkzcdc",
			},
			want: "cdc",
		},
		{
			name: "qbixwllojhsaeaehmsrk",
			args: args{
				input: "qbixwllojhsaeaehmsrk",
			},
			want: "aea",
		},
		{
			name: "abcdefghijklm",
			args: args{
				input: "abcccdeeeeefghiiiijklm",
			},
			want: "eeeee",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, longestPalindrome(tt.args.input))
		})
	}
}

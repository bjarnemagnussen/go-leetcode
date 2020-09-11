package problem0009

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input int
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example1",
			args: args{
				input: 121,
			},
			want: true,
		},
		{
			name: "example2",
			args: args{
				input: -121,
			},
			want: false,
		},
		{
			name: "example3",
			args: args{
				input: 10,
			},
			want: false,
		},
		{
			name: "0",
			args: args{
				input: 0,
			},
			want: true,
		},
		{
			name: "1",
			args: args{
				input: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, isPalindrome(tt.args.input))
		})
	}
}

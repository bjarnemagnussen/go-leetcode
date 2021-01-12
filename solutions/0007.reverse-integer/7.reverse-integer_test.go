package problem0007

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input int
}

func Test_compareVersion(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				input: 123,
			},
			want: 321,
		},
		{
			name: "example2",
			args: args{
				input: -123,
			},
			want: -321,
		},
		{
			name: "example3",
			args: args{
				input: 120,
			},
			want: 21,
		},
		{
			name: "zero",
			args: args{
				input: 0,
			},
			want: 0,
		},
		{
			name: "overflow1",
			args: args{
				input: 9463847412,
			},
			want: 0,
		},
		{
			name: "non-overflow1",
			args: args{
				input: 7463847412,
			},
			want: 2147483647,
		},
		{
			name: "overflow2",
			args: args{
				input: -9463847412,
			},
			want: 0,
		},
		{
			name: "non-overflow2",
			args: args{
				input: -8463847412,
			},
			want: -2147483648,
		},
		{
			name: "overflows-within",
			args: args{
				input: 8085774586302733229,
			},
			want: 0,
		},
		{
			name: "overflow3",
			args: args{
				input: 1534236469,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, reverse(tt.args.input))
		})
	}
}

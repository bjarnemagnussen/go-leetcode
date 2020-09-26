package problem0011

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input []int
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			name: "example2",
			args: args{
				input: []int{3, 9, 3, 4, 7, 2, 12, 6},
			},
			want: 45,
		},
		{
			name: "example3",
			args: args{
				input: []int{3, 9, 3, 4, 12, 2, 7, 6},
			},
			want: 36,
		},
		{
			name: "example4",
			args: args{
				input: []int{0, 0},
			},
			want: 0,
		},
		{
			name: "example5",
			args: args{
				input: []int{1, 1},
			},
			want: 1,
		},
		{
			name: "example6",
			args: args{
				input: []int{2, 3, 10, 5, 7, 8, 9},
			},
			want: 36,
		},
		{
			name: "example7",
			args: args{
				input: []int{1, 2, 4, 3},
			},
			want: 4,
		},
		{
			name: "example8",
			args: args{
				input: []int{2, 3, 4, 5, 18, 17, 6},
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, maxArea(tt.args.input))
		})
	}
}

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}

package problem0674

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input []int
}

func Test_findLengthOfLCIS(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				input: []int{1, 3, 5, 4, 7},
			},
			want: 3,
		},
		{
			name: "example2",
			args: args{
				input: []int{2, 2, 2, 2, 2},
			},
			want: 1,
		},
		{
			name: "increasing",
			args: args{
				input: []int{1, 2, 3, 4, 5},
			},
			want: 5,
		},
		{
			name: "decreasing",
			args: args{
				input: []int{5, 4, 3, 2, 1},
			},
			want: 1,
		},
		{
			name: "empty",
			args: args{
				input: []int{},
			},
			want: 0,
		},
		{
			name: "single",
			args: args{
				input: []int{1},
			},
			want: 1,
		},
		{
			name: "split1",
			args: args{
				input: []int{5, 6, 7, 8, 1, 2, 3},
			},
			want: 4,
		},
		{
			name: "split2",
			args: args{
				input: []int{5, 6, 7, 1, 2, 3, 4},
			},
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, findLengthOfLCIS(tt.args.input))
		})
	}
}

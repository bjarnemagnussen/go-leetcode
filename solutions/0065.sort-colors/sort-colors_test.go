package problem0065

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input []int
}

func Test_sortColors(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example",
			args: args{
				input: []int{2, 0, 2, 1, 1, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "inverse",
			args: args{
				input: []int{2, 2, 1, 1, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "sorted",
			args: args{
				input: []int{0, 0, 1, 1, 2, 2},
			},
			want: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name: "empty",
			args: args{
				input: []int{},
			},
			want: []int{},
		},
		{
			name: "same",
			args: args{
				input: []int{1, 1, 1},
			},
			want: []int{1, 1, 1},
		},
		{
			name: "all-pairs",
			args: args{
				input: []int{0, 1, 2, 0, 2, 1, 1, 0, 2, 1, 2, 0, 2, 0, 1, 2, 1, 0},
			},
			want: []int{0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortColors(tt.args.input)
			assert.Equal(t, tt.want, tt.args.input)
		})
	}
}

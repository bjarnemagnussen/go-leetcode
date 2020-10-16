package problem0015

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input []int
}

func Test_threeSum(t *testing.T) {
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example1",
			args: args{
				input: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "example2",
			args: args{
				input: []int{},
			},
			want: [][]int{},
		},
		{
			name: "example3",
			args: args{
				input: []int{0},
			},
			want: [][]int{},
		},
		{
			name: "custom1",
			args: args{
				input: []int{1, -1, 0, -1, 1, 0},
			},
			want: [][]int{{-1, 0, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := threeSum(tt.args.input)
			sort.SliceStable(res, func(i, j int) bool {
				if res[i][0] < res[j][0] || (res[i][0] == res[j][0] && res[i][1] < res[j][1]) || (res[i][0] == res[j][0] && res[i][1] == res[j][1] && res[i][2] < res[j][2]) {
					return true
				}
				return false
			})
			assert.Equal(t, tt.want, res)
		})
	}
}

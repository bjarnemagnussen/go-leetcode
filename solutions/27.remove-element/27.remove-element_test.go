package problem0027

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	nums []int
	val  int
}

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "example2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
		{
			name: "own1",
			args: args{
				nums: []int{3, 3, 2, 3},
				val:  3,
			},
			want: 1,
		},
		{
			name: "own2",
			args: args{
				nums: []int{3, 3, 3, 3},
				val:  3,
			},
			want: 0,
		},
		{
			name: "own3",
			args: args{
				nums: []int{3, 1, 3, 1, 1},
				val:  3,
			},
			want: 3,
		},
		{
			name: "own4",
			args: args{
				nums: []int{},
				val:  0,
			},
			want: 0,
		},
		{
			name: "own5",
			args: args{
				nums: []int{0, 4, 4, 0, 4, 4, 4, 0, 2},
				val:  4,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			counter := make(map[int]int, 0)
			for i := 0; i < len(tt.args.nums); i++ {
				counter[tt.args.nums[i]]++
			}

			ret := removeElement(tt.args.nums, tt.args.val)

			assert.Equal(t, tt.want, ret)

			for i := 0; i < ret; i++ {
				if tt.args.nums[i] == tt.args.val || counter[tt.args.nums[i]] <= 0 {
					t.Errorf("expected no value '%v' left in array", tt.args.val)
				}
				counter[tt.args.nums[i]]--
			}
		})
	}
}

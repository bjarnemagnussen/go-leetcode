package problem0017

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input string
}

func Test_letterCombinations(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "example1",
			args: args{
				input: "23",
			},
			want: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name: "example2",
			args: args{
				input: "",
			},
			want: []string{},
		},
		{
			name: "example3",
			args: args{
				input: "2",
			},
			want: []string{"a", "b", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, letterCombinations(tt.args.input))
		})
	}
}

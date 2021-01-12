package problem0165

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type args struct {
	input []string
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
				input: []string{"0.1", "1.1"},
			},
			want: -1,
		},
		{
			name: "example2",
			args: args{
				input: []string{"1.0.1", "1"},
			},
			want: 1,
		},
		{
			name: "example3",
			args: args{
				input: []string{"7.5.2.4", "7.5.3"},
			},
			want: -1,
		},
		{
			name: "example4",
			args: args{
				input: []string{"1.01", "1.001"},
			},
			want: 0,
		},
		{
			name: "example5",
			args: args{
				input: []string{"1.0", "1.0.0"},
			},
			want: 0,
		},
		{
			name: "same",
			args: args{
				input: []string{"0.1.2", "0.1.2"},
			},
			want: 0,
		},
		{
			name: "smaller",
			args: args{
				input: []string{"1.0.0", "1.0.1"},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, compareVersion(
				tt.args.input[0], tt.args.input[1]))
		})
	}
}

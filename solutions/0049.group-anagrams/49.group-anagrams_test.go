package problem0049

import (
	"reflect"
	"testing"
)

type args struct {
	strs []string
}

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "example1",
			args: args{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			name: "example2",
			args: args{
				strs: []string{""},
			},
			want: [][]string{{""}},
		},
		{
			name: "example3",
			args: args{
				strs: []string{"a"},
			},
			want: [][]string{{"a"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotGroups := groupAnagrams(tt.args.strs)

			if len(gotGroups) != len(tt.want) {
				t.Fatalf("got slice of length %d, wanted %d", len(gotGroups), len(tt.want))
			}

			for _, g := range gotGroups {
				visitedGroups := make(map[string]int)
				for _, gg := range g {
					visitedGroups[gg]++
				}

				var exists bool

				for _, w := range tt.want {
					visitedWanted := make(map[string]int)
					for _, ww := range w {
						visitedWanted[ww]++
					}
					if reflect.DeepEqual(visitedGroups, visitedWanted) {
						exists = true
					}
				}
				if !exists {
					t.Fatalf("string slice '%v' should not exist", g)
				}
			}
		})
	}
}

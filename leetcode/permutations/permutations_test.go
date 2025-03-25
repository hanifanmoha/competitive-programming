package permutations

import (
	"reflect"
	"sort"
	"testing"
)

func Test_permute(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 1, 2},
				{3, 2, 1},
			},
		},
		{
			name: "test1",
			args: args{
				nums: []int{0, 1},
			},
			want: [][]int{
				{0, 1},
				{1, 0},
			},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1},
			},
			want: [][]int{
				{1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := permute(tt.args.nums)
			sort.Slice(tt.want, func(i, j int) bool {
				for k := range tt.want[i] {
					if tt.want[i][k] < tt.want[j][k] {
						return true
					}
					if tt.want[i][k] > tt.want[j][k] {
						return false
					}
				}
				return false
			})
			sort.Slice(got, func(i, j int) bool {
				for k := range got[i] {
					if got[i][k] < got[j][k] {
						return true
					}
					if got[i][k] > got[j][k] {
						return false
					}
				}
				return false
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute() = %v, want %v", got, tt.want)
			}
		})
	}
}

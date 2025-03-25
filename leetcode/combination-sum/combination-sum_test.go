package combinationsum

import (
	"reflect"
	"sort"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{
				{2, 2, 3},
				{7},
			},
		},
		{
			name: "test2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{
				{2, 2, 2, 2},
				{2, 3, 3},
				{3, 5},
			},
		},
		{
			name: "test3",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			want: [][]int{},
		},
		{
			name: "test4",
			args: args{
				candidates: []int{1},
				target:     1,
			},
			want: [][]int{
				{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.args.candidates, tt.args.target)
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
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

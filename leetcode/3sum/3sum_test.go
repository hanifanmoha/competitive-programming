package three_sum

import (
	"reflect"
	"testing"
)

func Test_threeSum(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want [][]int
	}{
		{
			name: "TC1",
			args: []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "TC2",
			args: []int{0, 1, 1},
			want: [][]int{},
		},
		{
			name: "TC3",
			args: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "TC4",
			args: []int{-1, 0, 1, 0},
			want: [][]int{
				{-1, 0, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

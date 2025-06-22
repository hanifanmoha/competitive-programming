package findkpairswithsmallestsums

import (
	"reflect"
	"testing"
)

func Test_kSmallestPairs(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
		k     int
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
				nums1: []int{1, 7, 11},
				nums2: []int{2, 4, 6},
				k:     3,
			},
			want: [][]int{{1, 2}, {1, 4}, {1, 6}},
		},
		{
			name: "test2",
			args: args{
				nums1: []int{1, 1, 2},
				nums2: []int{1, 2, 3},
				k:     2,
			},
			want: [][]int{{1, 1}, {1, 1}},
		},
		{
			name: "test3",
			args: args{
				nums1: []int{1, 2, 4, 5, 6},
				nums2: []int{3, 5, 7, 9},
				k:     20,
			},
			want: [][]int{{1, 3}, {2, 3}, {1, 5}, {2, 5}, {4, 3}, {1, 7}, {5, 3}, {2, 7}, {4, 5}, {6, 3}, {1, 9}, {5, 5}, {2, 9}, {4, 7}, {6, 5}, {5, 7}, {4, 9}, {6, 7}, {5, 9}, {6, 9}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := kSmallestPairs(tt.args.nums1, tt.args.nums2, tt.args.k)
			if !reflect.DeepEqual(getSums(got), getSums(tt.want)) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getSums(arr [][]int) []int {
	res := make([]int, len(arr))
	for _, a := range arr {
		res = append(res, a[0]+a[1])
	}
	return res
}

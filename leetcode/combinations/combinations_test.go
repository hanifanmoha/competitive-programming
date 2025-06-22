package combinations

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_combine(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
		{
			name: "combinations of 4 and 2",
			args: args{n: 4, k: 2},
			want: [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}},
		},
		{
			name: "combinations of 4 and 3",
			args: args{n: 4, k: 3},
			want: [][]int{{1, 2, 3}, {1, 2, 4}, {1, 3, 4}, {2, 3, 4}},
		},
		{
			name: "combinations of 5 and 4",
			args: args{n: 5, k: 4},
			want: [][]int{{1, 2, 3, 4}, {1, 2, 3, 5}, {1, 2, 4, 5}, {1, 3, 4, 5}, {2, 3, 4, 5}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combine(tt.args.n, tt.args.k)
			want := tt.want

			sort.Slice(got, func(i, j int) bool {
				// concatenate the arrays
				gi := fmt.Sprint(got[i])
				gj := fmt.Sprint(got[j])
				return gi < gj
			})

			sort.Slice(want, func(i, j int) bool {
				// concatenate the arrays
				wi := fmt.Sprint(want[i])
				wj := fmt.Sprint(want[j])
				return wi < wj
			})

			fmt.Println("got", got)
			fmt.Println("want", want)

			if !reflect.DeepEqual(got, want) {
				t.Errorf("combine() = %v, want %v", got, tt.want)
			}
		})
	}
}

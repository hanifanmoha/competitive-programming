package ipo

import "testing"

func Test_findMaximizedCapital(t *testing.T) {
	type args struct {
		k       int
		w       int
		profits []int
		capital []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Example 1",
			args: args{
				k:       2,
				w:       0,
				profits: []int{1, 2, 3},
				capital: []int{0, 1, 1},
			},
			want: 4,
		},
		{
			name: "Example 2",
			args: args{
				k:       3,
				w:       0,
				profits: []int{1, 2, 3},
				capital: []int{0, 1, 2},
			},
			want: 6,
		},
		{
			name: "Example 3",
			args: args{
				k:       1,
				w:       0,
				profits: []int{1, 2, 3},
				capital: []int{1, 1, 2},
			},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{
				k:       1,
				w:       2,
				profits: []int{1, 2, 3},
				capital: []int{1, 1, 2},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaximizedCapital(tt.args.k, tt.args.w, tt.args.profits, tt.args.capital); got != tt.want {
				t.Errorf("findMaximizedCapital() = %v, want %v", got, tt.want)
			}
		})
	}
}

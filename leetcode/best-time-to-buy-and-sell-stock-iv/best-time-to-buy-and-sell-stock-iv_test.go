package besttimetobuyandsellstockiv

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		k      int
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Single transaction - increasing prices",
			args: args{
				k:      1,
				prices: []int{1, 2, 3, 4, 5},
			},
			want: 4,
		},
		{
			name: "Single transaction - decreasing prices",
			args: args{
				k:      1,
				prices: []int{5, 4, 3, 2, 1},
			},
			want: 0,
		},
		{
			name: "Single transaction - valley and peak",
			args: args{
				k:      1,
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "Single transaction - all same prices",
			args: args{
				k:      1,
				prices: []int{3, 3, 3, 3, 3},
			},
			want: 0,
		},
		{
			name: "Single transaction - two prices",
			args: args{
				k:      1,
				prices: []int{1, 10},
			},
			want: 9,
		},
		{
			name: "Example 1",
			args: args{
				k:      2,
				prices: []int{2, 4, 1},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				k:      2,
				prices: []int{3, 2, 6, 5, 0, 3},
			},
			want: 7,
		},
		{
			name: "Example 3",
			args: args{
				k:      2,
				prices: []int{2, 1, 4, 5, 2, 9, 7},
			},
			want: 11,
		},
		{
			name: "Three transactions - multiple peaks and valleys",
			args: args{
				k:      3,
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			},
			want: 22,
		},
		{
			name: "Three transactions - alternating ups and downs",
			args: args{
				k:      3,
				prices: []int{3, 8, 5, 1, 7, 8, 2, 6, 9, 4, 5, 3},
			},
			want: 19,
		},
		{
			name: "Three transactions - all increasing",
			args: args{
				k:      3,
				prices: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			},
			want: 11,
		},
		{
			name: "Three transactions - all decreasing",
			args: args{
				k:      3,
				prices: []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
			},
			want: 0,
		},
		{
			name: "Five transactions - long prices array",
			args: args{
				k:      5,
				prices: []int{3, 2, 6, 5, 0, 3, 8, 1, 4, 7, 2, 5, 9, 0, 6, 4, 8, 3, 7, 2, 5, 10, 1},
			},
			want: 37,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.k, tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

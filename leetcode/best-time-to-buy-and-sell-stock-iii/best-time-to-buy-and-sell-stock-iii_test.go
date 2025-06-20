package besttimetobuyandsellstockiii

import "testing"

func Test_maxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Example 1",
			args: args{prices: []int{3, 3, 5, 0, 0, 3, 1, 4}},
			want: 6,
		},
		{
			name: "Example 2",
			args: args{prices: []int{1, 2, 3, 4, 5}},
			want: 4,
		},
		{
			name: "Example 3",
			args: args{prices: []int{7, 6, 4, 3, 1}},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{prices: []int{2, 1, 2, 0, 1}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

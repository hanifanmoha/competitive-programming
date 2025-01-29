package main

import "testing"

func Test_canCompleteCircuit(t *testing.T) {
	type args struct {
		gas  []int
		cost []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "TC1",
			args: args{
				gas:  []int{1, 2, 3, 4, 5},
				cost: []int{3, 4, 5, 1, 2},
			},
			want: 3,
		},
		{
			name: "TC2",
			args: args{
				gas:  []int{2, 3, 4},
				cost: []int{3, 4, 3},
			},
			want: -1,
		},
		{
			name: "TC3",
			args: args{
				gas:  []int{1, 2, 3, 4, 5, 5, 70},
				cost: []int{2, 3, 4, 3, 9, 6, 2},
			},
			want: 6,
		},
		{
			name: "TC4",
			args: args{
				gas:  []int{2, 0, 1, 2, 3, 4, 0},
				cost: []int{0, 1, 0, 0, 0, 0, 11},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canCompleteCircuit(tt.args.gas, tt.args.cost); got != tt.want {
				t.Errorf("canCompleteCircuit() = %v, want %v", got, tt.want)
			}
		})
	}
}

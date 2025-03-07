package houserobber

import "testing"

func Test_rob(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test_case_1",
			args: args{nums: []int{1, 2, 3, 1}},
			want: 4,
		},
		{
			name: "test_case_2",
			args: args{nums: []int{2, 7, 9, 3, 1}},
			want: 12,
		},
		{
			name: "test_case_3",
			args: args{nums: []int{2, 1, 1, 2}},
			want: 4,
		},
		{
			name: "test_case_4",
			args: args{nums: []int{5, 5, 10, 100, 10, 5}},
			want: 110,
		},
		{
			name: "test_case_5",
			args: args{nums: []int{0}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.args.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}

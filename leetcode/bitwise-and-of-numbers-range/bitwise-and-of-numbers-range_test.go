package bitwiseandofnumbersrange

import "testing"

func Test_rangeBitwiseAnd(t *testing.T) {
	type args struct {
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.

		{
			name: "Example1",
			args: args{left: 5, right: 7},
			want: 4,
		},
		{
			name: "Example2",
			args: args{left: 0, right: 0},
			want: 0,
		},
		{
			name: "Example3",
			args: args{left: 1, right: 2147483647},
			want: 0,
		},
		{
			name: "SameNumber",
			args: args{left: 8, right: 8},
			want: 8,
		},
		{
			name: "SmallRange",
			args: args{left: 6, right: 7},
			want: 6,
		},
		{
			name: "PowerOfTwo",
			args: args{left: 16, right: 31},
			want: 16,
		},
		{
			name: "ZeroToOne",
			args: args{left: 0, right: 1},
			want: 0,
		},
		{
			name: "LargeRange",
			args: args{left: 123, right: 456},
			want: 0,
		},
		{
			name: "EdgeCaseMaxInt",
			args: args{left: 2147483646, right: 2147483647},
			want: 2147483646,
		},
		{
			name: "MiddleRange",
			args: args{left: 12, right: 15},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rangeBitwiseAnd(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("rangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

package basiccalculator

import "testing"

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{s: "1 + 1"},
			want: 2,
		},
		{
			name: "test2",
			args: args{s: " 2-1 + 2 "},
			want: 3,
		},
		{
			name: "test2.5",
			args: args{s: " 3 + 8 / 2 * 5"},
			want: 23,
		},
		{
			name: "test2.6",
			args: args{s: " 3 * 8 + 20 / 5"},
			want: 28,
		},
		{
			name: "test2.7",
			args: args{s: "((2+3)*5-(3*2))/3"},
			want: 6,
		},
		{
			name: "test2.8",
			args: args{s: "((10+2)*3-(4/2))/2"},
			want: 17,
		},
		{
			name: "test2.9",
			args: args{s: "((15-5)*(2+3))/5"},
			want: 10,
		},
		{
			name: "test2.10",
			args: args{s: "((8/2)+(3*4))-7"},
			want: 9,
		},
		{
			name: "test3",
			args: args{s: "(1+(4+5+2)-3)+(6+8)"},
			want: 23,
		},
		{
			name: "test4",
			args: args{s: "2-(5-6)"},
			want: 3,
		},
		{
			name: "test5",
			args: args{s: "3+2*2"},
			want: 7,
		},
		{
			name: "test6",
			args: args{s: "-2+ 1"},
			want: -1,
		},
		{
			name: "test8",
			args: args{s: "- (3 + ( - 4 - 1 + (-(5))))"},
			want: 7,
		},
		{
			name: "test7",
			args: args{s: "- (3 + ( - 4 + 5))"},
			want: -4,
		},
		{
			name: "test9",
			args: args{s: "1 - (-2)"},
			want: 3,
		},
		{
			name: "test10",
			args: args{s: "-3 - 2"},
			want: -5,
		},
		{
			name: "test11",
			args: args{s: "4 + (-2)"},
			want: 2,
		},
		{
			name: "test12",
			args: args{s: "-(2 + 3)"},
			want: -5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculate(tt.args.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}

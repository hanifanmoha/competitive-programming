package evaluate_reverse_polish_notation

import "testing"

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{tokens: []string{"2", "1", "+", "3", "*"}},
			want: 9,
		},
		{
			name: "test2",
			args: args{tokens: []string{"4", "13", "5", "/", "+"}},
			want: 6,
		},
		{
			name: "test3",
			args: args{tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
			want: 22,
		},
		{
			name: "test4",
			args: args{tokens: []string{"3", "4", "+", "2", "*", "7", "/"}},
			want: 2,
		},
		{
			name: "test5",
			args: args{tokens: []string{"5", "1", "2", "+", "4", "*", "+", "3", "-"}},
			want: 14,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := evalRPN(tt.args.tokens); got != tt.want {
				t.Errorf("evalRPN() = %v, want %v", got, tt.want)
			}
		})
	}
}

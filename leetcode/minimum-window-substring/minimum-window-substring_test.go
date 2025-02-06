package minimum_window_substring

import "testing"

func Test_minWindow(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "example0",
			args: args{
				s: "ab",
				t: "A",
			},
			want: "",
		},
		{
			name: "example1",
			args: args{
				s: "ADOBECODEBANC",
				t: "ABC",
			},
			want: "BANC",
		},
		{
			name: "example2",
			args: args{
				s: "a",
				t: "a",
			},
			want: "a",
		},
		{
			name: "example3",
			args: args{
				s: "a",
				t: "aa",
			},
			want: "",
		},
		{
			name: "edge case 1",
			args: args{
				s: "",
				t: "a",
			},
			want: "",
		},
		{
			name: "edge case 2",
			args: args{
				s: "a",
				t: "",
			},
			want: "",
		},
		{
			name: "edge case 3",
			args: args{
				s: "a",
				t: "b",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minWindow(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("minWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}

package interleavingstring

import "testing"

func Test_isInterleave(t *testing.T) {
	type args struct {
		s1 string
		s2 string
		s3 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "Example 1",
			args: args{s1: "aabcc", s2: "dbbca", s3: "aadbbcbcac"},
			want: true,
		},
		{
			name: "Example 2",
			args: args{s1: "aabcc", s2: "dbbca", s3: "aadbbbaccc"},
			want: false,
		},
		{
			name: "Both empty strings",
			args: args{s1: "", s2: "", s3: ""},
			want: true,
		},
		{
			name: "One empty, valid interleave",
			args: args{s1: "abc", s2: "", s3: "abc"},
			want: true,
		},
		{
			name: "One empty, invalid interleave",
			args: args{s1: "abc", s2: "", s3: "ab"},
			want: false,
		},
		{
			name: "Different lengths",
			args: args{s1: "a", s2: "b", s3: "abc"},
			want: false,
		},
		{
			name: "Interleaving with repeated chars",
			args: args{s1: "aa", s2: "ab", s3: "aaba"},
			want: true,
		},
		{
			name: "Not interleaving",
			args: args{s1: "abc", s2: "def", s3: "abdecf"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.args.s1, tt.args.s2, tt.args.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}

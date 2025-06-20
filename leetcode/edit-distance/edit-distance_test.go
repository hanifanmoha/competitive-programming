package editdistance

import "testing"

func Test_minDistance(t *testing.T) {
	type args struct {
		word1 string
		word2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "insertion only",
			args: args{
				word1: "abc",
				word2: "abucd",
			},
			want: 2,
		},
		{
			name: "deletion only",
			args: args{
				word1: "abcd",
				word2: "abc",
			},
			want: 1,
		},
		{
			name: "example1",
			args: args{
				word1: "horse",
				word2: "ros",
			},
			want: 3,
		},
		{
			name: "example2",
			args: args{
				word1: "intention",
				word2: "execution",
			},
			want: 5,
		},
		{
			name: "empty strings",
			args: args{
				word1: "",
				word2: "",
			},
			want: 0,
		},
		{
			name: "one empty, one non-empty",
			args: args{
				word1: "",
				word2: "abc",
			},
			want: 3,
		},
		{
			name: "identical strings",
			args: args{
				word1: "abc",
				word2: "abc",
			},
			want: 0,
		},
		{
			name: "completely different",
			args: args{
				word1: "abc",
				word2: "def",
			},
			want: 3,
		},
		{
			name: "single character difference",
			args: args{
				word1: "abc",
				word2: "ab",
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDistance(tt.args.word1, tt.args.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

package wordladder

import "testing"

func Test_ladderLength(t *testing.T) {
	type args struct {
		beginWord string
		endWord   string
		wordList  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Example 1",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log", "cog"},
			},
			want: 5,
		},
		{
			name: "Example 2 - no possible transformation",
			args: args{
				beginWord: "hit",
				endWord:   "cog",
				wordList:  []string{"hot", "dot", "dog", "lot", "log"},
			},
			want: 0,
		},
		{
			name: "Begin word equals end word",
			args: args{
				beginWord: "hit",
				endWord:   "hit",
				wordList:  []string{"hit"},
			},
			want: 1,
		},
		{
			name: "Single letter words",
			args: args{
				beginWord: "a",
				endWord:   "c",
				wordList:  []string{"a", "b", "c"},
			},
			want: 2,
		},
		{
			name: "No transformation needed",
			args: args{
				beginWord: "a",
				endWord:   "a",
				wordList:  []string{"a"},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ladderLength(tt.args.beginWord, tt.args.endWord, tt.args.wordList); got != tt.want {
				t.Errorf("ladderLength() = %v, want %v", got, tt.want)
			}
		})
	}
}

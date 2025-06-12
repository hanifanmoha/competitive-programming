package wordsearchii

import (
	"reflect"
	"testing"
)

func Test_findWords(t *testing.T) {
	type args struct {
		board [][]byte
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "Example 1",
			args: args{
				board: [][]byte{
					{'o', 'a', 'a', 'n'},
					{'e', 't', 'a', 'e'},
					{'i', 'h', 'k', 'r'},
					{'i', 'f', 'l', 'v'},
				},
				words: []string{"oath", "pea", "eat", "rain"},
			},
			want: []string{"oath", "eat"},
		},
		{
			name: "Example 2",
			args: args{
				board: [][]byte{
					{'a', 'b'},
					{'c', 'd'},
				},
				words: []string{"abcb"},
			},
			want: []string{},
		},
		{
			name: "Single letter board and word",
			args: args{
				board: [][]byte{
					{'a'},
				},
				words: []string{"a"},
			},
			want: []string{"a"},
		},
		{
			name: "No words found",
			args: args{
				board: [][]byte{
					{'a', 'b'},
					{'c', 'd'},
				},
				words: []string{"efg", "hij"},
			},
			want: []string{},
		},
		{
			name: "All words found",
			args: args{
				board: [][]byte{
					{'a', 'b'},
					{'c', 'd'},
				},
				words: []string{"ab", "abc", "abcd", "adcb"},
			},
			want: []string{"ab"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findWords(tt.args.board, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

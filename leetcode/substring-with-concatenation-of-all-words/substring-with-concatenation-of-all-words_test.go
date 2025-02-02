package substring_with_concatenation_of_all_words

import (
	"reflect"
	"testing"
)

func Test_findSubstring(t *testing.T) {
	type args struct {
		s     string
		words []string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.

		{
			name: "example1",
			args: args{
				s:     "barfoothefoobarman",
				words: []string{"foo", "bar"},
			},
			want: []int{0, 9},
		},
		{
			name: "example2",
			args: args{
				s:     "wordgoodgoodgoodbestword",
				words: []string{"word", "good", "best", "word"},
			},
			want: []int{},
		},
		{
			name: "edge case empty string",
			args: args{
				s:     "",
				words: []string{"foo", "bar"},
			},
			want: []int{},
		},
		{
			name: "edge case empty words",
			args: args{
				s:     "barfoothefoobarman",
				words: []string{},
			},
			want: []int{},
		},
		{
			name: "longer input",
			args: args{
				s:     "foobarbazbarfoobazbarfoo",
				words: []string{"foo", "bar", "baz"},
			},
			want: []int{0, 6, 9, 12, 15},
		},
		{
			name: "example3",
			args: args{
				s:     "barfoofoobarthefoobarman",
				words: []string{"bar", "foo", "the"},
			},
			want: []int{6, 9, 12},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findSubstring(tt.args.s, tt.args.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

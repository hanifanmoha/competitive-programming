package text_justification

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_fullJustify(t *testing.T) {
	input1 := []string{"This", "is", "an", "example", "of", "text", "justification."}
	input2 := []string{"What", "must", "be", "acknowledgment", "shall", "be"}
	input3 := []string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"}

	output1 := []string{
		"This    is    an",
		"example  of text",
		"justification.  ",
	}

	output2 := []string{
		"What   must   be",
		"acknowledgment  ",
		"shall be        ",
	}

	output3 := []string{
		"Science  is  what we",
		"understand      well",
		"enough to explain to",
		"a  computer.  Art is",
		"everything  else  we",
		"do                  ",
	}

	type args struct {
		words    []string
		maxWidth int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TC1",
			args: args{
				words:    input1,
				maxWidth: 16,
			},
			want: output1,
		},
		{
			name: "TC2",
			args: args{
				words:    input2,
				maxWidth: 16,
			},
			want: output2,
		},
		{
			name: "TC3",
			args: args{
				words:    input3,
				maxWidth: 20,
			},
			want: output3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fullJustify(tt.args.words, tt.args.maxWidth); !reflect.DeepEqual(got, tt.want) {
				// t.Errorf("fullJustify() = %v, want %v", got, tt.want)
				t.Error(tt.name)
				fmt.Println("=====WANT=====")
				for _, s := range tt.want {
					fmt.Println(s)
				}
				fmt.Println("=====GOTT=====")
				for _, s := range got {
					fmt.Println(s)
				}
				fmt.Println("==============")
			}
		})
	}
}

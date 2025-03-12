package lettercombinationsofaphonenumber

import (
	"reflect"
	"sort"
	"testing"
)

func Test_letterCombinations(t *testing.T) {
	type args struct {
		digits string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
		{
			name: "test0",
			args: args{
				digits: "23",
			},
			want: []string{
				"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf",
			},
		},

		{
			name: "test1",
			args: args{
				digits: "234",
			},
			want: []string{
				"adg", "adh", "adi", "aeg", "aeh", "aei", "afg", "afh", "afi",
				"bdg", "bdh", "bdi", "beg", "beh", "bei", "bfg", "bfh", "bfi",
				"cdg", "cdh", "cdi", "ceg", "ceh", "cei", "cfg", "cfh", "cfi",
			},
		},
	}
	for _, tt := range tests {
		sort.Strings(tt.want)

		t.Run(tt.name, func(t *testing.T) {
			got := letterCombinations(tt.args.digits)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("letterCombinations() = %v, want %v", got, tt.want)
			}
		})
	}
}

package minimumgeneticmutation

import "testing"

func Test_minMutation(t *testing.T) {
	type args struct {
		startGene string
		endGene   string
		bank      []string
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
				startGene: "AACCGGTT",
				endGene:   "AACCGGTA",
				bank:      []string{"AACCGGTA"},
			},
			want: 1,
		},
		{
			name: "Example 2",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AAACGGTA",
				bank:      []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"},
			},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{
				startGene: "AAAAACCC",
				endGene:   "AACCCCCC",
				bank:      []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			},
			want: 3,
		},
		{
			name: "No possible mutation",
			args: args{
				startGene: "AAAAACCC",
				endGene:   "AACCCCCC",
				bank:      []string{"AAAACCCC", "AAACCCCC"},
			},
			want: -1,
		},
		{
			name: "Start equals end",
			args: args{
				startGene: "AACCGGTT",
				endGene:   "AACCGGTT",
				bank:      []string{"AACCGGTT"},
			},
			want: 0,
		},
		{
			name: "Complex path with dead end",
			args: args{
				startGene: "AAAAAAAA",
				endGene:   "CCCCCCCC",
				bank:      []string{"AAAAAAAA", "AAAAAAAC", "AAAAAACC", "AAAAACCC", "AAAACCCC", "AACACCCC", "ACCACCCC", "ACCCCCCC", "CCCCCCCA"},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMutation(tt.args.startGene, tt.args.endGene, tt.args.bank); got != tt.want {
				t.Errorf("minMutation() = %v, want %v", got, tt.want)
			}
		})
	}
}

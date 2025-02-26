package surroundedregions

import "testing"

func Test_solve(t *testing.T) {
	type args struct {
		board [][]byte
	}
	tests := []struct {
		name     string
		args     args
		expected [][]byte
	}{
		{
			name: "example1",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'O', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "example2",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X'},
					{'X', 'O', 'X', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "example3",
			args: args{
				board: [][]byte{
					{'X', 'O', 'X', 'X'},
					{'X', 'O', 'X', 'X'},
					{'X', 'X', 'O', 'X'},
					{'X', 'O', 'X', 'X'},
				},
			},
			expected: [][]byte{
				{'X', 'O', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			},
		},
		{
			name: "example4",
			args: args{
				board: [][]byte{
					{'X', 'X', 'X', 'X', 'X'},
					{'O', 'O', 'X', 'O', 'X'},
					{'X', 'O', 'O', 'X', 'X'},
				},
			},
			expected: [][]byte{
				{'X', 'X', 'X', 'X', 'X'},
				{'O', 'O', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X', 'X'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
			for i := range tt.args.board {
				for j := range tt.args.board[i] {
					if tt.args.board[i][j] != tt.expected[i][j] {
						t.Errorf("solve() = %v, want %v", tt.args.board, tt.expected)
					}
				}
			}
		})
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solve(tt.args.board)
		})
	}
}

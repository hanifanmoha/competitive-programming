package spiral_matrix

import (
	"reflect"
	"testing"
)

func Test_spiralOrder(t *testing.T) {
	// Matrix 1: 3x3
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	output1 := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}

	// Matrix 2: 3x4
	matrix2 := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}

	output2 := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}

	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "TC1",
			args: args{
				matrix: matrix1,
			},
			want: output1,
		},
		{
			name: "TC2",
			args: args{
				matrix: matrix2,
			},
			want: output2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// fmt.Println("INPUT")
			// fmt.Println(len(tt.args.matrix), len(tt.args.matrix[0]))
			// for row := range tt.args.matrix {
			// 	fmt.Println(row)
			// }
			// fmt.Println("======")
			if got := spiralOrder(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("spiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

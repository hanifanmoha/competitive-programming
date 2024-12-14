package main

import (
	"reflect"
	"testing"
)

func Test_gameOfLife(t *testing.T) {

	input1 := [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}

	// Expected Output 1
	output1 := [][]int{
		{0, 0, 0},
		{1, 0, 1},
		{0, 1, 1},
		{0, 1, 0},
	}

	// Input 2
	input2 := [][]int{
		{1, 1},
		{1, 0},
	}

	// Expected Output 2
	output2 := [][]int{
		{1, 1},
		{1, 1},
	}

	tests := []struct {
		name   string
		input  [][]int
		output [][]int
	}{
		// TODO: Add test cases.
		{
			name:   "TC1",
			input:  input1,
			output: output1,
		},
		{
			name:   "TC2",
			input:  input2,
			output: output2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.input)
			if !reflect.DeepEqual(tt.input, tt.output) {
				t.Errorf("gameOfLife() = %v, want %v", tt.input, tt.output)
			}
		})
	}
}

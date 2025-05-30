package evaluatedivision

import (
	"reflect"
	"testing"
)

func Test_calcEquation(t *testing.T) {
	type args struct {
		equations [][]string
		values    []float64
		queries   [][]string
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		// TODO: Add test cases.
		{
			name: "basic division",
			args: args{
				equations: [][]string{{"a", "b"}, {"b", "c"}},
				values:    []float64{2.0, 3.0},
				queries:   [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			},
			want: []float64{6.0, 0.5, -1.0, 1.0, -1.0},
		},
		{
			name: "self division and unknowns",
			args: args{
				equations: [][]string{{"x", "y"}},
				values:    []float64{4.0},
				queries:   [][]string{{"x", "x"}, {"y", "y"}, {"x", "z"}, {"z", "z"}},
			},
			want: []float64{1.0, 1.0, -1.0, -1.0},
		},
		{
			name: "cycle in equations",
			args: args{
				equations: [][]string{{"a", "b"}, {"b", "c"}, {"c", "a"}},
				values:    []float64{2.0, 3.0, 0.5},
				queries:   [][]string{{"a", "c"}, {"c", "b"}, {"b", "a"}},
			},
			want: []float64{6.0, 0.3333333333333333, 0.5},
		},
		{
			name: "disconnected graphs",
			args: args{
				equations: [][]string{{"a", "b"}, {"c", "d"}},
				values:    []float64{2.0, 3.0},
				queries:   [][]string{{"a", "d"}, {"b", "a"}, {"c", "d"}},
			},
			want: []float64{-1.0, 0.5, 3.0},
		},
		{
			name: "complex graph with multiple variables",
			args: args{
				equations: [][]string{{"x1", "x4"}, {"x2", "x3"}, {"x1", "x2"}, {"x2", "x5"}},
				values:    []float64{3.0, 0.5, 3.4, 5.6},
				queries:   [][]string{{"x1", "x5"}, {"x4", "x5"}, {"x1", "x3"}, {"x5", "x5"}, {"x5", "x1"}, {"x3", "x4"}, {"x4", "x3"}, {"x6", "x6"}, {"x0", "x0"}},
				// queries: [][]string{{"x1", "x5"}, {"x4", "x5"}},
			},
			want: []float64{19.04, 6.34667, 1.7, 1.0, 0.05252, 1.76471, 0.56667, -1.0, -1.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcEquation(tt.args.equations, tt.args.values, tt.args.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcEquation() = %v, want %v", got, tt.want)
			}
		})
	}
}

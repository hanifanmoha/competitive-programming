package factorialtrailingzeroes

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "n=3",
			args: args{n: 3},
			want: 0,
		},
		{
			name: "n=5",
			args: args{n: 5},
			want: 1,
		},
		{
			name: "n=0",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "n=10",
			args: args{n: 10},
			want: 2,
		},
		{
			name: "n=25",
			args: args{n: 25},
			want: 6,
		},
		{
			name: "n=100",
			args: args{n: 100},
			want: 24,
		},
		{
			name: "n=125",
			args: args{n: 125},
			want: 31,
		},
		{
			name: "n=1000",
			args: args{n: 1000},
			want: 249,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}

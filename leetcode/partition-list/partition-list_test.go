package partition_list

import (
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "example1",
			args: args{
				head: &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{5, &ListNode{2, nil}}}}}},
				x:    3,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{2, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}}},
		},
		{
			name: "example2",
			args: args{
				head: &ListNode{2, &ListNode{1, nil}},
				x:    2,
			},
			want: &ListNode{1, &ListNode{2, nil}},
		},
		{
			name: "empty list",
			args: args{
				head: nil,
				x:    1,
			},
			want: nil,
		},
		{
			name: "all nodes less than x",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{2, nil}}},
				x:    3,
			},
			want: &ListNode{1, &ListNode{2, &ListNode{2, nil}}},
		},
		{
			name: "all nodes greater than or equal to x",
			args: args{
				head: &ListNode{4, &ListNode{5, &ListNode{6, nil}}},
				x:    3,
			},
			want: &ListNode{4, &ListNode{5, &ListNode{6, nil}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}

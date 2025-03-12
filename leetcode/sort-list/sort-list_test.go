package sortlist

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{
			name: "example 1",
			args: args{
				head: &ListNode{-1, &ListNode{5, &ListNode{3, &ListNode{4, &ListNode{0, nil}}}}},
			},
			want: &ListNode{-1, &ListNode{0, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			name: "example 2",
			args: args{
				head: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, &ListNode{6, nil}}}}}},
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
		},
		{
			name: "example 3",
			args: args{
				head: &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}}},
			},
			want: &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{1, nil}}}}},
		},
		{
			name: "example 4",
			args: args{
				head: &ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}},
			},
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			a1 := list2arr(sortList(tt.args.head))
			a2 := list2arr(tt.want)

			fmt.Println("RESULT", a1)
			fmt.Println("EXPECTED", a2)

			if !reflect.DeepEqual(a1, a2) {
				t.Errorf("sortList() = %v, want %v", a1, a2)
			}
		})
	}
}

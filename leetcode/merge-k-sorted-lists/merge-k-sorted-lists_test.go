package merge_k_sorted_lists

import (
	"fmt"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
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
				lists: []*ListNode{
					{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}},
					{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}},
					{Val: 2, Next: &ListNode{Val: 6}},
				},
			},
			want: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}},
			},
		},
		{
			name: "empty lists",
			args: args{
				lists: []*ListNode{},
			},
			want: nil,
		},
		{
			name: "all lists empty",
			args: args{
				lists: []*ListNode{nil, nil, nil},
			},
			want: nil,
		},
		{
			name: "single list",
			args: args{
				lists: []*ListNode{
					{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
				},
			},
			want: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}},
			},
		},
		{
			name: "lists with one element each",
			args: args{
				lists: []*ListNode{
					{Val: 1},
					{Val: 0},
					{Val: 2},
				},
			},
			want: &ListNode{
				Val:  0,
				Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//
			p1 := tt.want
			p2 := mergeKLists(tt.args.lists)

			fmt.Println("Want:", list2Arr(p1))
			fmt.Println("Got:", list2Arr(p2))

			for p1 != nil || p2 != nil {
				if p1 == nil || p2 == nil || p1.Val != p2.Val {
					t.Errorf("mergeKLists() = %v, want %v", p2, p1)
					break
				}
				p1 = p1.Next
				p2 = p2.Next

			}

			// if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			// }
		})
	}
}

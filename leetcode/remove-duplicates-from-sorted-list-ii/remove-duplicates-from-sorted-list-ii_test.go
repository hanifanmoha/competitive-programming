package remove_duplicates_from_sorted_list_ii

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
			name: "Test case 1",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, nil}}}}}}},
			},
			want: &ListNode{1, &ListNode{2, &ListNode{5, nil}}},
		},
		{
			name: "Test case 2",
			args: args{
				head: &ListNode{1, &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}},
			},
			want: &ListNode{2, &ListNode{3, nil}},
		},
		{
			name: "Test case 3",
			args: args{
				head: &ListNode{1, &ListNode{2, &ListNode{2, nil}}},
			},
			want: &ListNode{1, nil},
		},
		{
			name: "Test case 4",
			args: args{
				head: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{2, nil}}}},
			},
			want: nil,
		},
		{
			name: "Test case 5",
			args: args{
				head: nil,
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

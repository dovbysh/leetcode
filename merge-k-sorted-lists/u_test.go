package merge_k_sorted_lists

import (
	"testing"
)

func TestListNode_String(t *testing.T) {
	tests := []struct {
		name string
		l    *ListNode
		want string
	}{
		{
			name: "1->4->5",
			l: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
			want: "1->4->5",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.l.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_arrayToListNode(t *testing.T) {
	tests := []struct {
		name string
		a    []int
		want string
	}{
		{
			name: "1->2->3->4",
			a:    []int{1, 2, 3, 4},
			want: "1->2->3->4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayToListNode(tt.a).String(); got != tt.want {
				t.Errorf("arrayToListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

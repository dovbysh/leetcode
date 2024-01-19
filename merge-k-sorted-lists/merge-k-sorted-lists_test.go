package merge_k_sorted_lists

import (
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1->1->2->3->4->4->5->6",
			args: args{
				arrayToListNodeArray([][]int{
					{1, 4, 5}, {1, 3, 4}, {2, 6},
				}),
			},
			want: "1->1->2->3->4->4->5->6",
		},
		{
			name: "[[],[1]]",
			args: args{
				arrayToListNodeArray([][]int{
					{}, {1},
				}),
			},
			want: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := mergeKLists(tt.args.lists)
			if got.String() != tt.want {
				t.Errorf("mergeKLists() = %v, want %v", got.String(), tt.want)
			}
		})
	}
}

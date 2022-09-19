package add_two_numbers

import (
	"reflect"
	"testing"
)

func slice2List(z []int) *ListNode {
	var first *ListNode
	prev := first
	for _, v := range z {
		current := ListNode{Val: v}
		if first == nil {
			first = &current
			prev = first
			continue
		}
		prev.Next = &current
		prev = &current
	}
	return first
}

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "342 + 465 = 807",
			args: args{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val:  3,
							Next: nil,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val:  4,
							Next: nil,
						},
					},
				},
			},
			want: []int{7, 0, 8},
		},
		{
			name: "9999999",
			args: args{
				l1: slice2List([]int{9, 9, 9, 9, 9, 9, 9}),
				l2: slice2List([]int{9, 9, 9, 9}),
			},
			want: []int{8, 9, 9, 9, 0, 0, 0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addTwoNumbers(tt.args.l1, tt.args.l2)
			gotSlice := make([]int, 0, len(tt.want))
			c := got
			for c != nil {
				gotSlice = append(gotSlice, c.Val)
				c = c.Next
			}
			if !reflect.DeepEqual(gotSlice, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", gotSlice, tt.want)
			}
		})
	}
}

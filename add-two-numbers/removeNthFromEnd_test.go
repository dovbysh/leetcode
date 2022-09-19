package add_two_numbers

import (
	"reflect"
	"testing"
)

func list2Slice(f *ListNode) []int {
	var r []int
	c := f
	for c != nil {
		r = append(r, c.Val)
		c = c.Next
	}
	return r
}

func Test_removeNthFromEnd(t *testing.T) {
	type args struct {
		head []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "[1,2,3,4,5], n = 2",
			args: args{
				head: []int{1, 2, 3, 4, 5},
				n:    2,
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "[1], n = 1",
			args: args{
				head: []int{1},
				n:    1,
			},
			want: nil,
		},
		{
			name: "[1,2], n = 1",
			args: args{
				head: []int{1, 2},
				n:    1,
			},
			want: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := list2Slice(removeNthFromEnd(slice2List(tt.args.head), tt.args.n))
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}

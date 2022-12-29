package median_of_two_sorted_arrays

import "testing"

func Test_findMedianSortedArrays(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1,3 - 2",
			args: args{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2,
		},
		{
			name: "1,2,3",
			args: args{
				nums1: []int{1},
				nums2: []int{2, 3},
			},
			want: 2,
		},
		{
			name: "1,2,3,4",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{4},
			},
			want: 2.5,
		},
		{
			name: "1,2,3,4,5,6,7",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{4, 5, 6, 7},
			},
			want: 4,
		},
		{
			name: "1,2,3,4,5,6,7,8",
			args: args{
				nums1: []int{1, 2, 3},
				nums2: []int{4, 5, 6, 7, 8},
			},
			want: 4.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.args.nums1, tt.args.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}

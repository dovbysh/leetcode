package serialize_and_deserialize_binary_tree

import (
	"testing"
)

func Test_CodecSerDeser(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want *TreeNode
	}{
		{
			name: "nil",
			root: nil,
			want: nil,
		},
		{
			name: "one",
			root: &TreeNode{
				Val: 1,
			},
			want: &TreeNode{
				Val: 1,
			},
		},
		{
			name: "three",
			root: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
			want: &TreeNode{
				Val:   3,
				Left:  &TreeNode{Val: 4},
				Right: &TreeNode{Val: 5},
			},
		},
		{
			name: "ex",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
			want: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
		},
	}

	ser := Constructor()
	deser := Constructor()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := ser.serialize(tt.root)
			got := deser.deserialize(data)
			if !tt.root.Eq(got) {
				t.Errorf("%s fail", tt.name)
			}
		})
	}
}

func TestTreeNode_Deep(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want int
	}{
		{
			name: "3",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
			want: 3,
		},
		{
			name: "4",
			root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 4},
					Right: &TreeNode{
						Val:   5,
						Right: &TreeNode{Val: 0},
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Deep(tt.root); got != tt.want {
				t.Errorf("Deep() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_serialize(t *testing.T) {
	tests := []struct {
		name string
		root *TreeNode
		want string
	}{
		{
			name: "",
			root: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 8},
					},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 6},
				},
			},
			want: "-1:l:0,0:l:1,1:l:3,2:r:8,1:r:4,0:r:2,5:l:5,5:r:6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &Codec{}
			if got := this.serialize(tt.root); got != tt.want {
				t.Errorf("serialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCodec_deserialize(t *testing.T) {
	tests := []struct {
		data string
		want *TreeNode
	}{
		{
			data: "-1:l:0,0:l:1,1:l:3,2:r:8,1:r:4,0:r:2,5:l:5,5:r:6",
			want: &TreeNode{
				Val: 0,
				Left: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val:   3,
						Right: &TreeNode{Val: 8},
					},
					Right: &TreeNode{Val: 4},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 5},
					Right: &TreeNode{Val: 6},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.data, func(t *testing.T) {
			this := &Codec{}
			got := this.deserialize(tt.data)
			if !got.Eq(tt.want) {
				t.Errorf("deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

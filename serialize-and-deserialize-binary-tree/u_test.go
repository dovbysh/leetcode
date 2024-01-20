package serialize_and_deserialize_binary_tree

import (
	"strings"
	"testing"
)

func TestNothing(t *testing.T) {
	ss := make([]string, 3)
	ss[1] = "2"
	zs := strings.Join(ss, ",")
	sss := strings.Split(zs, ",")
	t.Log(zs)
	t.Log(sss)
}
func TestTreeNode_Eq(t *testing.T) {
	type fields struct {
		Val   int
		Left  *TreeNode
		Right *TreeNode
	}
	type args struct {
		d *TreeNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "z",
			fields: fields{
				Val:   1,
				Left:  nil,
				Right: nil,
			},
			args: args{},
			want: false,
		},
		{
			name: "x",
			fields: fields{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
			args: args{d: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			}},
			want: true,
		},
		{
			name: "x",
			fields: fields{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
			},
			args: args{d: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
				Right: &TreeNode{
					Val:   3,
					Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 0}},
					Right: &TreeNode{Val: 5},
				},
			}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TreeNode{
				Val:   tt.fields.Val,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := s.Eq(tt.args.d); got != tt.want {
				t.Errorf("Eq() = %v, want %v", got, tt.want)
			}
		})
	}
}

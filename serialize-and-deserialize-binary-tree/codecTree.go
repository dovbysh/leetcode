// https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

package serialize_and_deserialize_binary_tree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

func ser(n *TreeNode, ss []string, i int, l string) []string {
	if n == nil {
		return nil
	}
	sss := append(ss, strconv.Itoa(i)+":"+l+":"+strconv.Itoa(n.Val))
	parentI := len(sss) - 1
	if n.Left != nil {
		sss = ser(n.Left, sss, parentI, "l")
	}
	if n.Right != nil {
		sss = ser(n.Right, sss, parentI, "r")
	}
	return sss
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	ss := make([]string, 0)
	ss = ser(root, ss, -1, "l")
	return strings.Join(ss, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	ss := strings.Split(data, ",")
	if ss[0] == "" {
		return nil
	}
	root := &TreeNode{}
	parents := make([]*TreeNode, len(ss))
	parents[0] = root
	for i := 0; i < len(ss); i++ {
		deser(ss[i], parents, i)
	}
	return root
}

func deser(ss string, parents []*TreeNode, i int) {
	s := strings.Split(ss, ":")
	parentI, err := strconv.Atoi(s[0])
	if err != nil {
		return
	}
	v, err := strconv.Atoi(s[2])
	if err != nil {
		return
	}
	if parentI == -1 {
		parents[0].Val = v
		return
	}
	n := &TreeNode{Val: v}
	if s[1] == "l" {
		parents[parentI].Left = n
	} else {
		parents[parentI].Right = n
	}
	parents[i] = n
}

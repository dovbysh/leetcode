package serialize_and_deserialize_binary_tree

func (s *TreeNode) Eq(d *TreeNode) bool {
	if s == nil && d == nil {
		return true
	}
	if s == nil || d == nil {
		return false
	}

	return s.Val == d.Val && s.Left.Eq(d.Left) && s.Right.Eq(d.Right)
}

func Deep(s *TreeNode) int {
	if s == nil {
		return 0
	}
	if s.Left == nil && s.Right == nil {
		return 1
	}

	return 1 + max(Deep(s.Left), Deep(s.Right))
}

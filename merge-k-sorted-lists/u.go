package merge_k_sorted_lists

import "strconv"

// [[1,4,5],[1,3,4],[2,6]]

func arrayToListNodeArray(a [][]int) []*ListNode {
	res := make([]*ListNode, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = arrayToListNode(a[i])
	}
	return res
}
func arrayToListNode(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}
	first := &ListNode{
		Val:  a[0],
		Next: nil,
	}
	cur := first
	for i := 1; i < len(a); i++ {
		cur.Next = &ListNode{Val: a[i]}
		cur = cur.Next
	}
	return first
}

func (l *ListNode) String() string {
	if l == nil {
		return ""
	}
	if l.Next == nil {
		return strconv.Itoa(l.Val)
	}
	return strconv.Itoa(l.Val) + "->" + l.Next.String()
}

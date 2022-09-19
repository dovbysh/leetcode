package add_two_numbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var first, prev *ListNode
	e1, e2 := l1, l2
	var pp int
	for true {
		if e1 == nil && e2 == nil {
			break
		}
		current := ListNode{}
		if e1 != nil {
			current.Val += e1.Val
			e1 = e1.Next
		}
		if e2 != nil {
			current.Val += e2.Val
			e2 = e2.Next
		}
		current.Val += pp
		pp = 0
		if current.Val > 9 {
			pp = 1
			current.Val = current.Val % 10
		}
		if first == nil {
			first = &current
		}
		if prev != nil {
			prev.Next = &current
		}
		prev = &current
	}
	if pp > 0 {
		prev.Next = &ListNode{Val: pp}
	}
	return first
}

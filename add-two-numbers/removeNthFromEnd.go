package add_two_numbers

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	z := make([]*ListNode, 0, n)
	c := head
	for c != nil {
		z = append(z, c)
		c = c.Next
	}
	if len(z) == 1 {
		return nil
	}
	if len(z) == n {
		return z[0].Next
	} else {
		nexti := len(z) - n + 1
		if nexti == len(z) {
			z[len(z)-1-n].Next = nil
		} else {
			z[len(z)-1-n].Next = z[nexti]
		}
	}
	return head
}

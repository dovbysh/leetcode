package merge_k_sorted_lists

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var minI int
	c := make([]*ListNode, len(lists))
	copy(c, lists)
	var first, cur *ListNode
	if c[minI] == nil {
		for k := 0; k < len(c); k++ {
			if c[k] != nil {
				minI = k
				break
			}
		}
	}
	for c[minI] != nil {
		for i := 0; i < len(c); i++ {
			if c[i] == nil {
				continue
			}
			if c[minI] == nil {
				minI = i
				continue
			}
			if c[minI].Val > c[i].Val {
				minI = i
			}
		}
		if c[minI] == nil {
			break
		}
		if first == nil {
			first = c[minI]
			cur = first
		} else {
			cur.Next = c[minI]
			cur = cur.Next
		}
		c[minI] = c[minI].Next
		if c[minI] == nil {
			for k := 0; k < len(c); k++ {
				if c[k] != nil {
					minI = k
					break
				}
			}
		}
	}
	return first
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    	if head == nil {
		return head
	}

	first := &ListNode{}

	result := &first.Next

	node := head
	for node != nil {
		if node.Val == val {
			node = node.Next
		} else {
			addNode := &ListNode{
				Val: node.Val,
			}

			first.Next = addNode
			first = first.Next
			node = node.Next
		}
	}

	return *result
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
    if list1 == nil && list2 == nil {
		return nil
	}
    
    head, firstStart := makeHead(list1, list2)

	firstNode := list1
	secondNode := list2

	if firstStart {
		firstNode = firstNode.Next
	} else {
		secondNode = secondNode.Next
	}

	result := head

	for firstNode != nil || secondNode != nil {
		addNode, firstSmall := makeHead(firstNode, secondNode)

		head.Next = addNode
		head = head.Next

		if firstSmall {
			firstNode = firstNode.Next
		} else {
			secondNode = secondNode.Next
		}
	}

	return result
}


func makeHead(list1 *ListNode, list2 *ListNode) (*ListNode, bool) {

	if list1 == nil {
		return &ListNode{Val: list2.Val}, false
	}

	if list2 == nil {
		return &ListNode{Val: list1.Val}, true
	}

	if list1.Val > list2.Val {
		return &ListNode{Val: list2.Val}, false
	} else {
		return &ListNode{Val: list1.Val}, true
	}


}

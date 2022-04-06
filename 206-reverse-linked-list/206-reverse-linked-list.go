/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    	var currentNode *ListNode
	for head != nil {
		tempNode := &ListNode{Val: head.Val}
		tempNode.Next = currentNode
		currentNode = tempNode
		head = head.Next
	}
	return currentNode
}
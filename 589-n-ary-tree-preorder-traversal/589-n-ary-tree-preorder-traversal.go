/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    result := make([]int, 0)
    
   if root == nil {
		return result
	}

	return getChild(root)
}

func getChild(node *Node) []int {

	arr := make([]int, 0)
	arr = append(arr, node.Val)

	for _, child := range node.Children {
		arr = append(arr, getChild(child)...)
	}

	return arr
}

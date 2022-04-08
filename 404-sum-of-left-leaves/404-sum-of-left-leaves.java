/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public int sumOfLeftLeaves(TreeNode root) {
         return getSumOfLeftLeaves(root, false);
    }
    private int getSumOfLeftLeaves(TreeNode treeNode, boolean isLeft) {

        if (treeNode == null) {
            return 0;
        }

        if (treeNode.left == null && treeNode.right == null && isLeft) {
            return treeNode.val;
        }

        return getSumOfLeftLeaves(treeNode.left,true) + getSumOfLeftLeaves(treeNode.right, false);
    }
}
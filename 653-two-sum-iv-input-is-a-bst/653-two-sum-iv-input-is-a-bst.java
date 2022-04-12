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
    public boolean findTarget(TreeNode root, int k) {
        Set<Integer> integerSet = getDepthSum(root,new HashSet<>());

        for (Integer integer : integerSet) {
      if (integerSet.contains(k-integer) && k-integer != integer) {
                return true;
            }
        }

        return false;
    }
    
    public Set<Integer> getDepthSum(TreeNode root, Set<Integer> integerSet) {
        if (root == null) {
            return integerSet;
        }

        integerSet.add(root.val);

        getDepthSum(root.left,integerSet);
        getDepthSum(root.right,integerSet);

        return integerSet;
    }
}
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
    public List<List<Integer>> levelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList<>();

        Map<Integer, List<TreeNode>> maps = getIntList(new HashMap<>(),root, 0);


        IntStream.range(0,maps.size()).forEach(e-> {
            List<TreeNode> treeNodes = maps.get(e);
            if (treeNodes == null) {
                return;
            }
            List<Integer> depthList = treeNodes.stream().map(treeNode -> treeNode.val).collect(Collectors.toList());

            result.add(depthList);
        });


        return result;
    }
        private Map<Integer, List<TreeNode>> getIntList(Map<Integer, List<TreeNode>> result ,TreeNode root, int depth) {

        if (root ==null) {
            return result;
        }

        List<TreeNode> treeNodes = result.get(depth)== null ? new ArrayList<>() : result.get(depth);
        treeNodes.add(root);

        result.put(depth, treeNodes);

       getIntList(result, root.left, depth+1);
       getIntList(result, root.right, depth+1);

       return result;
    }
}
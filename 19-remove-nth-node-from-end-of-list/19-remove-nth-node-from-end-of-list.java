/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class Solution {
    public ListNode removeNthFromEnd(ListNode head, int n) {
        Map<Integer, ListNode> listNodeList =new HashMap<>();
        ListNode result = new ListNode(0, head);

        int count = 0;
        listNodeList.put(count, result);

        while (head != null) {
            count++;
            listNodeList.put(count, head);
            head= head.next;
        }

        ListNode targetBeforeNode = listNodeList.get(count-n);
        targetBeforeNode.next = listNodeList.get(count-n+2);
        
        return result.next;
    }
}
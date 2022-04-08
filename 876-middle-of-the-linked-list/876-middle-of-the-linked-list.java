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
    public ListNode middleNode(ListNode head) {
            Map<Integer, ListNode> listNodeMap = new HashMap();
        
        int count = 0;
        int point = 0;

        while (head != null) {
            count++;
            listNodeMap.put(count, head);
            head = head.next;
        }


        point = count/2 +1;
        
        return listNodeMap.get(point);
    }
}
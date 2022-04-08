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
    public int getDecimalValue(ListNode head) {
                List<Integer> integerList = integerList(head);

        int mulVal = (int) Math.pow(2, integerList.size()-1);

        int result =0;

        for (Integer integer : integerList) {
            result += mulVal* integer;
            mulVal /=2;
        }

        return result;
    }
        private List<Integer> integerList(ListNode head) {
        List<Integer> integerList = new ArrayList<>();

        while (head != null) {
            integerList.add(head.val);
            head = head.next;
        }

        return integerList;
    }
}
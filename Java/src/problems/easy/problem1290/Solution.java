package problems.easy.problem1290;

class Solution {
    public int getDecimalValue(ListNode head) {
        var result = 0;

        while(head != null) {
            if (head.val == 0) {
                result <<= 1;
            } else {
                result <<= 1;
                result++;
            }

            head = head.next;
        }

        return result;
    }
}

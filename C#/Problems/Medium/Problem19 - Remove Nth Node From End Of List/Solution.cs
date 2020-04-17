using LeetCode.Portfolio;

namespace LeetCode.Problems.Medium.Problem19
{
    class Solution
    {
        public ListNode RemoveNthFromEnd(int[] array, int n)
        {
            ListNode node = this.generateListNode(array);
            return this.RemoveNthFromEnd(node, n);
        }

        public ListNode RemoveNthFromEnd(ListNode head, int n)
        {
            if (head.next == null && n == 1)
                return null;

            var position = this.goToEnd(head, head.next, n);

            return position == 0 ? head.next : head;
        }


        public string getValue(ListNode node)
        {
            return string.Join("", getValueFromListNode(node));
        }


        private int goToEnd(ListNode current, ListNode nextNode, int n)
        {
            if (current.next == null)
                return --n;

            var position = this.goToEnd(nextNode, nextNode.next, n);

            if (position == 0)
            {
                current.next = nextNode.next;
            }

            return --position;
        }

        private string getValueFromListNode(ListNode listNode)
        {
            string value = listNode.val.ToString();
            if (listNode.next != null)
            {
                value += getValueFromListNode(listNode.next);
            }

            return value;
        }

        private ListNode generateListNode(int[] array)
        {
            ListNode listNode = null;

            for (int i = array.Length - 1; 0 <= i; i--)
            {
                var item = array[i];

                if (listNode == null)
                {
                    listNode = new ListNode(item);
                    continue;
                }

                ListNode node = new ListNode(item);
                node.next = listNode;

                listNode = node;
            }

            return listNode;
        }
    }
}

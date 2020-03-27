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
            return null;
        }


        public string getValue(ListNode node)
        {
            return string.Join("", getValueFromListNode(node));
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

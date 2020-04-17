using LeetCode.Portfolio;

namespace LeetCode.Problems.Easy.Problem876
{
    class Solution
    {
        private int _middle = 0;

        public ListNode MiddleNode(int[] arr)
        {
            var head = this.generateListNode(arr);
            return this.MiddleNode(head);
        }

        public ListNode MiddleNode(ListNode head)
        {
            if (head.next == null)
                return head;

            return this.checkNext(head, 1);
        }

        private ListNode checkNext(ListNode node, int count)
        {
            if (node.next == null)
            {
                _middle = (count / 2) + 1;
                return node;
            }
            var returned = checkNext(node.next, count + 1);

            if (_middle == count)
                return node;
            else
                return returned;
            
        }

        public string getValue(ListNode node)
        {
            return string.Join("", getValueFromListNode(node));
        }

        private void addNode(int val, ref ListNode mainNode, ref ListNode temp)
        {
            if (mainNode == null)
            {
                mainNode = new ListNode(val);
            }
            else
            {
                if (temp == null)
                {
                    temp = new ListNode(val);
                    mainNode.next = temp;
                }
                else
                {
                    ListNode node = new ListNode(val);
                    temp.next = node;
                    temp = node;
                }
            }
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

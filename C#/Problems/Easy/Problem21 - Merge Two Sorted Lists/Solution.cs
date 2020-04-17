using LeetCode.Portfolio;

namespace LeetCode.Problems.Easy.Problem21
{
    class Solution
    {
        public ListNode MergeTwoLists(int[] array1, int[] array2)
        {
            var node1 = this.generateListNode(array1);
            var node2 = this.generateListNode(array2);

            return this.MergeTwoLists(node1, node2);
        }
        public ListNode MergeTwoLists(ListNode l1, ListNode l2)
        {
            ListNode mainNode = null;
            ListNode temp = null;

            while(l1 != null && l2 != null)
            {
                if (l1.val <= l2.val)
                {
                    this.addNode(l1.val, ref mainNode, ref temp);
                    l1 = l1.next;
                }
                else
                {
                    this.addNode(l2.val, ref mainNode, ref temp);
                    l2 = l2.next;
                }
            }

            while (l1 != null)
            {
                this.addNode(l1.val, ref mainNode, ref temp);
                l1 = l1.next;
            }

            while (l2 != null)
            {
                this.addNode(l2.val, ref mainNode, ref temp);
                l2 = l2.next;
            }

            return mainNode;
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

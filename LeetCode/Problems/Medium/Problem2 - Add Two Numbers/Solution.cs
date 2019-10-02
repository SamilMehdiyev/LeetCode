using System.Linq;
using System.Numerics;

namespace LeetCode.Problems.Medium.Problem2
{
    public class ListNode
    {
        public int val;
        public ListNode next;
        public ListNode(int x) { val = x; }
    }

    class Solution
    {
        public ListNode AddTwoNumbers(BigInteger num1, BigInteger num2)
        {
            return AddTwoNumbers(generateListNode(num1), generateListNode(num2));
        }

        public ListNode AddTwoNumbers(ListNode l1, ListNode l2)
        {
            string strOne = getValueFromListNode(l1);
            string strTwo = getValueFromListNode(l2);

            var numOne = parseValueToNumber(strOne);
            var numTwo = parseValueToNumber(strTwo);

            return generateListNode(numOne + numTwo);
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

        private BigInteger parseValueToNumber(string str)
        {
            return BigInteger.Parse(string.Join("", str.Reverse()));
        }

        private ListNode generateListNode(BigInteger number)
        {
            ListNode listNode = null;

            foreach (var item in number.ToString())
            {
                if(listNode == null)
                {
                    listNode = new ListNode((int)char.GetNumericValue(item));
                    continue;
                }

                ListNode node = new ListNode((int)char.GetNumericValue(item));
                node.next = listNode;

                listNode = node;
            }

            return listNode;
        }
    }
}

using System;
using LeetCode.Portfolio;

namespace LeetCode.Problems.Easy.Problem543
{
    public class Solution
    {
        public int DiameterOfBinaryTree(TreeNode root)
        {
            int[] longestPath = new int[1];
            processTree(root, longestPath);

            return longestPath[0];
        }

        private int processTree(TreeNode node, int[] longestPath)
        {
            if (node == null)
                return 0;

            var leftResult = processTree(node.left, longestPath);
            var rightResult = processTree(node.right, longestPath);

            longestPath[0] = Math.Max(longestPath[0], leftResult + rightResult);
            return Math.Max(leftResult, rightResult) + 1;
        }
    }
}

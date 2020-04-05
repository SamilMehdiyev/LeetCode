namespace LeetCode.Problems.Easy.Problem53
{
    class Solution
    {
        public int MaxSubArray(int[] nums)
        {
            if (nums == null || nums.Length == 0)
                return 0;

            var currentSum = 0;
            var maxSum = 0;

            for (int i = 1; i < nums.Length; i++)
            {
                currentSum = this.findMax(nums[i], currentSum + nums[i]);
                if (currentSum > maxSum)
                    maxSum = currentSum;
            }

            return maxSum;
        }

        private int findMax(int first, int second)
        {
            return first >= second ? first : second;
        }
    }
}

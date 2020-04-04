namespace LeetCode.Problems.Easy.Problem53
{
    class Solution
    {
        public int MaxSubArray(int[] nums)
        {
            if (nums == null || nums.Length == 0)
                return 0;

            //if (nums.Length == 1)
            //    return nums[0];

            var sum = 0;
            for (int i = 0; i < nums.Length; i++)
            {
                sum += nums[i];
            }

            var start = 0;
            var end = nums.Length - 1;

            var maxSum = sum;

            while (start < end)
            {
                if (nums[start] < nums[end])
                {
                    start += 1;
                    sum -= nums[start - 1];
                }
                else
                {
                    end -= 1;
                    sum -= nums[end + 1];
                }


                if (maxSum <= sum)
                    maxSum = sum;

            }

            return maxSum;
        }
    }
}

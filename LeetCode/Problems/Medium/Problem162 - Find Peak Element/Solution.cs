namespace LeetCode.Problems.Medium.Problem162
{
    class Solution
    {
        public int FindPeakElement(int[] nums)
        {
            var start = 0;
            var end = nums.Length;

            while (start < end)
            {
                var mid = (start + end) / 2;

                var next = (mid + 1) % nums.Length;
                var prev = (mid + nums.Length - 1) % nums.Length;

                if (nums[prev] <= nums[mid] && nums[mid] >= nums[next])
                    return mid;

                if (nums[prev] <= nums[mid] && nums[mid] <= nums[next])
                    start = mid + 1;
                else
                    end = mid;
            }

            return -1;
        }
    }
}

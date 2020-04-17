namespace LeetCode.Problems.Medium.Problem153
{
    class Solution
    {
        public int FindMin(int[] nums)
        {

            if (nums == null || nums.Length == 0)
                return -1;

            if (nums.Length == 1)
                return nums[0];

            var start = 0;
            var end = nums.Length - 1;

            while (start < end)
            {
                var mid = (start + end) / 2;
                
                if(nums[start] <= nums[end])
                {
                    return nums[mid];
                }

                if (nums[mid] > nums[end])
                    start = mid + 1;
                else
                    end = mid;
            }

            return nums[start];
        }
    }
}

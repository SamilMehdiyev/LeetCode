namespace LeetCode.Problems.Medium.Problem34
{
    class Solution
    {
        public int[] SearchRange(int[] nums, int target)
        {
            int[] result = new int[2];
            result[0] = findTargetStartIndex(nums, 0, nums.Length - 1, target);
            result[1] = findTargetEndIndex(nums, 0, nums.Length - 1, target);
            return result;
        }

        private int findTargetStartIndex(int[] nums, int start, int end, int target)
        {
            int idx = -1;
            while (start <= end)
            {
                int mid = (start + end) / 2;

                if (nums[mid] >= target)
                    end = mid - 1;
                else
                    start = mid + 1;

                if (nums[mid] == target) idx = mid;
            }
            return idx;
        }

        private int findTargetEndIndex(int[] nums, int start, int end, int target)
        {
            int idx = -1;
            while (start <= end)
            {
                int mid = (start + end) / 2;

                if (nums[mid] <= target)
                    start = mid + 1;
                else
                    end = mid - 1;

                if (nums[mid] == target) idx = mid;
            }
            return idx;
        }
    }
}

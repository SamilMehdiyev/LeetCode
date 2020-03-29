namespace LeetCode.Problems.Easy.Problem35
{
    class Solution
    {
        public int SearchInsert(int[] nums, int target)
        {
            var low = 0;
            var high = nums.Length - 1;

            while (low <= high)
            {
                var mid = (low + high) / 2;
                var element = nums[mid];

                if (element == target)
                    return mid;

                if (target < element)
                    high = mid - 1;
                else
                    low = mid + 1;
            }

            if (high < 0 || low == nums.Length)
                return low;

            return nums[low] > target ? low : high;
        }
    }
}

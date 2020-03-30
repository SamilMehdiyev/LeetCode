namespace LeetCode.Problems.Medium.Problem33
{
    class Solution
    {
        public int Search(int[] nums, int target)
        {
            if (nums == null || nums.Length == 0)
                return -1;

            var pivot = this.findPivot(nums);

            var start = 0;
            var end = nums.Length - 1;

            var result = this.search(nums, start, pivot, target);

            if (result != -1)
                return result;

            result = this.search(nums, pivot, end, target);

            if (result != -1)
                return result;

            return -1;
        }

        private int search(int[] nums, int start, int end, int target)
        {
            while (start <= end)
            {
                var mid = (start + end) / 2;
                var element = nums[mid];

                if (element == target)
                    return mid;

                if (element < target)
                    start = mid + 1;
                else
                    end = mid - 1;
            }

            return -1;
        }

        private int findPivot(int[] nums)
        {
            var start = 0;
            var end = nums.Length - 1;



            while (start <= end)
            {
                if (nums[start] <= nums[end])
                    return start;

                var mid = (start + end) / 2;

                var next = (mid + 1) % nums.Length;
                var prev = (mid + nums.Length - 1) % nums.Length;

                if (nums[mid] <= nums[mid + 1] && nums[mid] <= nums[mid - 1])
                    return mid;

                if (nums[mid] <= nums[end])
                    end = mid - 1;

                if (nums[mid] >= nums[start])
                    start = mid + 1;
            }

            return -1;
        }
    }
}

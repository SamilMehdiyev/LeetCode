namespace LeetCode.Problems.Easy.Problem283
{
    class Solution
    {
        public int[] MoveZeroes(int[] nums)
        {
            var start = 0;
            var end = nums.Length - 1;
            var index = 0;

            while (start < end && index < nums.Length)
            {
                if (nums[index] != 0)
                {
                    index++;
                    continue;
                }
                else
                {
                    var nonZeroIndex = this.findNextNonZeroIndex(nums, start, end);

                    if (nonZeroIndex == -1)
                    {
                        index++;
                        continue;
                    }

                    var temp = nums[start];
                    nums[start] = nums[nonZeroIndex];
                    nums[nonZeroIndex] = temp;

                    start++;
                    if (nums[end] == 0)
                        end--;
                }
            }

            return nums;
        }
        private int findNextNonZeroIndex(int[] nums, int start, int end)
        {
            var index = -1;

            for (int i = start; i <= end; i++)
            {
                if (nums[i] == 0)
                    continue;

                if (index == -1)
                {
                    index = i;
                    break;
                }
            }

            return index;
        }
    }
}

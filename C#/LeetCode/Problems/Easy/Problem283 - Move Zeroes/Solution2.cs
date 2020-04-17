namespace LeetCode.Problems.Easy.Problem283
{
    class Solution2
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
                    var minIndex = this.findMinimumValueIndex(nums, start, end);

                    if (minIndex == -1)
                    {
                        index++;
                        continue;
                    }

                    var temp = nums[start];
                    nums[start] = nums[minIndex];
                    nums[minIndex] = temp;

                    start++;
                    if(nums[end] == 0)
                        end--;
                }
            }

            return nums;
        }
        private int findMinimumValueIndex(int[] nums, int start, int end)
        {
            var minIndex = -1;

            for (int i = start; i <= end; i++)
            {
                if (nums[i] == 0)
                    continue;

                if (minIndex == -1)
                {
                    minIndex = i;
                    continue;
                }

                if (nums[i] < nums[minIndex])
                    minIndex = i;
            }

            return minIndex;
        }
    }
}

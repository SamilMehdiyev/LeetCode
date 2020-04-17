namespace LeetCode.Problems.Easy.Problem26
{
    class Solution
    {
        public int RemoveDuplicates(int[] nums)
        {
            var length = nums.Length;
            for (int i = 1; i < length; i++)
            {
                if (nums[i] == nums[i - 1])
                {
                    this.slideArray(ref nums, i);
                    i--;
                    length--;
                }
            }

            return length;
        }

        private void slideArray(ref int[] nums, int index)
        {
            for (int i = index + 1; i < nums.Length; i++)
            {
                nums[i - 1] = nums[i];
            }
        }
    }
}

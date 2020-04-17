namespace LeetCode.Problems.Easy.Problem27
{
    class Solution
    {
        public int RemoveElement(int[] nums, int val)
        {
            var length = nums.Length;
            for (int i = 0; i < length; i++)
            {
                if (nums[i] == val)
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

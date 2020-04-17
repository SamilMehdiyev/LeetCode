using System.Collections.Generic;

namespace LeetCode.Problems.Easy.Problem136
{
    class Solution
    {
        public int SingleNumber(int[] nums)
        {
            Dictionary<int, int> pairs = new Dictionary<int, int>();

            for (int i = 0; i < nums.Length; i++)
            {
                var element = nums[i];

                if (pairs.ContainsKey(element))
                    pairs[element] = pairs[element] + 1;
                else
                    pairs[element] = 1;
            }

            int result = 0;
            foreach (var item in pairs)
            {
                result = item.Key;

                if (item.Value == 1)
                    break;
            }

            return result;
        }
    }
}

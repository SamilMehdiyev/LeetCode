using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem26
{
    class Solution2
    {
        public int RemoveDuplicates(int[] nums)
        {
            if (nums == null)
                return 0;

            if (nums.Length < 2)
            {
                return nums.Length;
            }

            int i = 1;

            for (int j = 1; j < nums.Length; j++)
            {
                if (nums[j] != nums[j - 1])
                    nums[i++] = nums[j];
            }

            return i;
        }
    }
}

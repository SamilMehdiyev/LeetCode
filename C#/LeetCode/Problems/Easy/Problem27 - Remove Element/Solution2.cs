using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem27
{
    class Solution2
    {
        public int RemoveElement(int[] nums, int val)
        {
            if (nums == null)
                return 0;

            int i = 0;

            for (int j = 0; j < nums.Length; j++)
            {
                if (nums[j] != val)
                    nums[i++] = nums[j];
            }

            return i;
        }
    }
}

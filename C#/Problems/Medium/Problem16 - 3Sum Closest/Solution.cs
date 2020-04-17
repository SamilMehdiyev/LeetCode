using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Medium.Problem16
{
    class Solution
    {
        public int ThreeSumClosest(int[] nums, int target)
        {
            var result = nums[0] + nums[1] + nums[nums.Length - 1];

            Array.Sort(nums);

            for (int i = 0; i < nums.Length; i++)
            {
                if (i > 0 && nums[i] == nums[i - 1])
                {
                    continue;
                }

                var start = i + 1;
                var end = nums.Length - 1;

                while (start < end)
                {
                    var sum = nums[i] + nums[start] + nums[end];

                    if (sum < target)
                    {
                        start++;
                    }
                    else
                    {
                        end--;
                    }

                    if (Math.Abs(sum - target) < Math.Abs(result - target))
                    {
                        result = sum;
                    }
                }
            }

            return result;
        }
    }
}

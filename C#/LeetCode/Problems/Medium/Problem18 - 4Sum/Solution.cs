using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Medium.Problem18
{
    class Solution
    {
        IList<IList<int>> list = new List<IList<int>>();

        public IList<IList<int>> FourSum(int[] nums, int target)
        {
            Array.Sort(nums);

            for (int i = 0; i < nums.Length; i++)
            {
                if (i > 0 && nums[i] == nums[i - 1])
                {
                    continue;
                }

                for (int j = i + 1; j < nums.Length; j++)
                {
                    if (j > i + 1 && nums[j] == nums[j - 1])
                    {
                        continue;
                    }

                    var l = j + 1;
                    var k = nums.Length - 1;

                    while (l < k)
                    {
                        var sum = nums[i] + nums[j] + nums[l] + nums[k];

                        if (sum < target)
                        {
                            l++;
                        }

                        if (sum > target)
                        {
                            k--;
                        }

                        if (sum == target)
                        {
                            List<int> quadruplet = new List<int>();

                            quadruplet.Add(nums[i]);
                            quadruplet.Add(nums[j]);
                            quadruplet.Add(nums[l]);
                            quadruplet.Add(nums[k]);

                            list.Add(quadruplet);

                            l++;
                            k--;

                            while (l < k && nums[l] == nums[l - 1]) l++;  // skip same result
                            while (l < k && nums[k] == nums[k + 1]) k--;  // skip same result
                        }
                    }
                }
            }

            return list;
        }
    }
}

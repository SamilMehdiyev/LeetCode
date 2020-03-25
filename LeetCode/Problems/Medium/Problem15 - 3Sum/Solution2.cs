using System;
using System.Collections.Generic;

namespace LeetCode.Problems.Medium.Problem15
{
    class Solution2
    {
        IList<IList<int>> list = new List<IList<int>>();

        public IList<IList<int>> ThreeSum(int[] nums)
        {
            Array.Sort(nums);

            for (int i = 0; i < nums.Length; i++)
            {
                if (i > 0 && nums[i] == nums[i - 1])
                {
                    continue;
                }

                var j = i + 1;
                var k = nums.Length - 1;

                while (j < k)
                {
                    var sum = nums[i] + nums[j] + nums[k];

                    if (sum < 0)
                    {
                        j++;
                    }

                    if (sum > 0)
                    {
                        k--;
                    }

                    if (sum == 0)
                    {
                        List<int> triples = new List<int>();

                        triples.Add(nums[i]);
                        triples.Add(nums[j]);
                        triples.Add(nums[k]);

                        list.Add(triples);

                        j++;
                        k--;

                        while (j < k && nums[j] == nums[j - 1]) j++;  // skip same result
                        while (j < k && nums[k] == nums[k + 1]) k--;  // skip same result
                    }
                }
            }

            return list;
        }

        private bool existValues(int a, int b, int c)
        {
            foreach (var item in list)
            {
                if (
                    (item[0] == a && item[1] == b && item[2] == c) ||
                    (item[0] == a && item[2] == b && item[1] == c) ||
                    (item[1] == a && item[2] == b && item[0] == c) ||
                    (item[1] == a && item[0] == b && item[2] == c) ||
                    (item[2] == a && item[0] == b && item[1] == c) ||
                    (item[2] == a && item[1] == b && item[0] == c)
                   )
                    return true;
            }
            return false;
        }
    }
}

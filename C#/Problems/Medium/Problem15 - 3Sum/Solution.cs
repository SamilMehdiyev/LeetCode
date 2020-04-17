using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Medium.Problem15
{
    class Solution
    {
        IList<IList<int>> list = new List<IList<int>>();

        public IList<IList<int>> ThreeSum(int[] nums)
        {
            int x, y, z;

            for (x = 0; x < nums.Length; x++)
            {
                for (y = x + 1; y < nums.Length; y++)
                {
                    for (z = y + 1; z < nums.Length; z++)
                    {
                        if (nums[x] > 0 && nums[y] > 0 && nums[z] > 0)
                            continue;

                        if (nums[x] < 0 && nums[y] < 0 && nums[z] < 0)
                            continue;

                        if (nums[x] + nums[y] + nums[z] == 0 && !this.existValues(nums[x], nums[y], nums[z]))
                        {
                            List<int> triples = new List<int>();

                            triples.Add(nums[x]);
                            triples.Add(nums[y]);
                            triples.Add(nums[z]);

                            list.Add(triples);
                        }
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

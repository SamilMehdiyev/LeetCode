using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Medium.Problem11
{
    class Solution3
    {
        public int MaxArea(int[] height)
        {
            var max = 0;

            var begin = 0;
            var end = height.Length - 1;

            while (begin != end)
            {
                var x = end - begin;
                var y = height[end];

                if (height[begin] < height[end])
                {
                    y = height[begin];
                }

                var result = x * y;

                if (result > max)
                {
                    max = result;
                }


                if (height[begin] < height[end])
                {
                    begin++;
                    continue;
                }

                end--;
            }

            return max;
        }
    }
}

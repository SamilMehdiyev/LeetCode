using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Medium.Problem11
{
    class Solution2
    {
        public int MaxArea(int[] height)
        {
            var max = 0;

            for (int i = 0; i <= height.Length - 1; i++)
            {
                var x = 1;

                for (int j = i + 1; j <= height.Length - 1; j++)
                {
                    var y = height[i] < height[j] ? height[i] : height[j];

                    var result = y * x;

                    if (result > max)
                    {
                        max = result;
                    }
                    
                    x++;
                }
            }
            return max;
        }
    }
}

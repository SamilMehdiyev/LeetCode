using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem1046
{
    class Solution
    {
        public int LastStoneWeight(int[] stones)
        {
            var stoneCount = stones.Length;
            var weight = 0;

            while (stoneCount > 1)
            {
                var fhi = (stones[0] > stones[1]) ? 0 : 1;  // first heaviest index
                var shi = (stones[0] > stones[1]) ? 1 : 0;  // second heaviest index

                for (int i = 2; i < stoneCount; i++)
                {
                    if (stones[i] > stones[fhi])
                    {
                        shi = fhi;
                        fhi = i;
                        continue;
                    }
                    
                    if (stones[i] > stones[shi])
                    {
                        shi = i;
                        continue;
                    }
                }

                stones[fhi] -= stones[shi];

                for (int i = shi + 1; i < stoneCount; i++)
                {
                    stones[i - 1] = stones[i];
                }

                stoneCount--;

            }

            if (stones.Length > 0)
                weight = stones[0];

            return weight;
        }
    }
}

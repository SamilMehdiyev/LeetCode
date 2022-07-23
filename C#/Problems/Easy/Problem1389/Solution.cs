using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem1389
{
    public class Solution {
        public int[] CreateTargetArray(int[] nums, int[] index) {
            if (nums.Length == 0 || index.Length == 0) {
                return new int[]{};
            }
            
            var target = new List<int>();
            
            for (int i = 0; i < index.Length; i++) {
                var idx = index[i];
                var value = nums[i];
                
                target.Insert(idx, value);
            }
            
            return target.ToArray();
        }
    }
}
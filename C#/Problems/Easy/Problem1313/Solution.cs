using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem1313
{
    public class Solution {
        public int[] DecompressRLElist(int[] nums) {
            var result = new List<int>();
            
            for (int i = 0; i < nums.Length; i += 2) {
                var freq = nums[i];
                var val = nums[i+1];
                
                while (freq > 0) {
                    result.Add(val);
                    freq--;
                }
            }
            
            return result.ToArray();
        }
    }
}
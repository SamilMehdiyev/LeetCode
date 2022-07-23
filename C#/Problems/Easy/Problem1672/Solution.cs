using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem1672
{
    public class Solution {
        public int MaximumWealth(int[][] accounts) {
            if (accounts.Length == 0) { 
                return 0;
            }
            
            var maxWealth = 0;
            for (int i = 0;i < accounts.Length; i++){
                var total = 0;
                for (int j = 0; j < accounts[i].Length; j++) {
                    total += accounts[i][j];
                }
                if (total > maxWealth) {
                    maxWealth = total;
                }
            }
            return maxWealth;
        }
    }
}
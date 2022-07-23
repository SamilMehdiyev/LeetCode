using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem1773
{
    public class Solution {
        public int CountMatches(IList<IList<string>> items, string ruleKey, string ruleValue) {
            var matches = 0;
            
            foreach(var item in items) {
                if (ruleKey == "type" && ruleValue == item[0]) {
                    matches++;
                }
                
                if (ruleKey == "color" && ruleValue == item[1]) {
                    matches++;
                }
                
                if (ruleKey == "name" && ruleValue == item[2]) {
                    matches++;
                }
            }
            
            return matches;
        }
    }
}
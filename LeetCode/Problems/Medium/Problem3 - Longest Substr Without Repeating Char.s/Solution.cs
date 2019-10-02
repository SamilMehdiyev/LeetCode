using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Medium.Problem3
{
    class Solution
    {
        public int LengthOfLongestSubstring(string s)
        {
            StringBuilder substring = new StringBuilder();

            if (s.Length == 0)
                return 0;

            int maxLength = 0;
            int substringLength = 0;

            loop:
            foreach (var item in s)
            {
                if(!substring.ToString().Contains(item))
                {
                    substring.Append(item);
                    substringLength++;
                    continue;
                }

                if (substringLength > maxLength)
                {
                    maxLength = substringLength;
                }

                substring.Remove(0, substringLength);

                substringLength = 0;

                break;
            }

            s = s.Substring(1);

            if (s.Length > 0) goto loop;

            if (substringLength > maxLength)
            {
                maxLength = substringLength;
            }

            return maxLength;
        }
    }
}

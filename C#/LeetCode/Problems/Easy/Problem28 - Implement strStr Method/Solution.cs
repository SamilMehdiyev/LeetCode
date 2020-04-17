using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem28
{
    class Solution
    {
        public int StrStr(string haystack, string needle)
        {
            if (needle.Length == 0)
                return 0;
            if (haystack.Length == 0)
                return -1;

            var startPoint = -1;
            var index = 0;
            var length = needle.Length;

            for (int i = 0; i < haystack.Length && index < needle.Length; i++)
            {
                if (haystack[i] == needle[index])
                {
                    if (startPoint == -1)
                        startPoint = i;
                    index++;
                }
                else
                {
                    if (startPoint != -1)
                        i = i - index;

                    index = 0;
                    startPoint = -1;
                    continue;
                }
            }

            if (index < needle.Length)
                return -1;

            return startPoint;
        }
    }
}

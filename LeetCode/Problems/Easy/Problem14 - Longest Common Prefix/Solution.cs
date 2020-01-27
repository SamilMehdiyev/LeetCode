using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem14
{
    class Solution
    {
        public string LongestCommonPrefix(string[] strs)
        {
            if (strs.Length == 0)
                return "";

            var length = 0;

            var minLengthStr = findMinimumLengthString(strs);

            if (minLengthStr.Length == 0)
                return "";

            while (length < minLengthStr.Length)
            {
                for (int i = 0; i < strs.Length; i++)
                {
                    var item = strs[i];

                    if (item[length] != minLengthStr[length])
                    {
                        goto result;
                    }
                }

                length++;
            }


            result: return minLengthStr.Substring(0, length);
        }

        private string findMinimumLengthString(string[] strs)
        {
            var min = strs[0];

            for (int i = 1; i < strs.Length; i++)
            {
                if (strs[i].Length < min.Length)
                    min = strs[i];
            }

            return min;
        }
    }
}

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

            string result = strs[0];
            var dictionary = new Dictionary<int, char>();

            if (strs[0].Length == 0)
                return "";

            prepareDictionary(strs[0], dictionary);

            for (int i = 1; i < strs.Length; i++)
            {
                var item = strs[i];
                var commonPrefixExist = true;

                if (item.Length == 0)
                    return "";

                for (int j = 0; j < item.Length; j++)
                {
                    if (dictionary[j] == item[j])
                        continue;

                    result = item.Substring(0, j);

                    if (j == 0)
                        commonPrefixExist = false;

                    break;
                }

                if (commonPrefixExist)
                    continue;

                break;
            }

            return result;
        }

        private void prepareDictionary(string str, Dictionary<int, char> dictionary)
        {
            for (int i = 0; i < str.Length; i++)
            {
                dictionary[i] = str[i];
            }
        }
    }
}

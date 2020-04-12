using System.Collections.Generic;

namespace LeetCode.Problems.Easy.Problem1408
{
    class Solution
    {

        public IList<string> StringMatching(string[] words)
        {
            List<string> strs = new List<string>();

            for (int i = 0; i < words.Length; i++)
            {
                for (int j = 0; j < words.Length; j++)
                {
                    if (i == j)
                        continue;

                    var indexOf = words[i].IndexOf(words[j]);
                    if (indexOf == -1)
                        continue;

                    strs.Add(words[j]);
                }
            }

            return strs;
        }
    }
}

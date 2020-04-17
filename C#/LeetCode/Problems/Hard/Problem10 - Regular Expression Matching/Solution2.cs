using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Hard.Problem10
{
    class Solution2
    {
        public bool IsMatch(string s, string p)
        {
            string text = s;
            string pattern = p;

            if (string.IsNullOrEmpty(pattern)) return string.IsNullOrEmpty(text);

            bool match = (!string.IsNullOrEmpty(text) && (pattern[0] == text[0] || pattern[0] == '.'));

            if (pattern.Length >= 2 && pattern[1].Equals('*'))
            {
                return (IsMatch(text, pattern.Substring(2)) || (match && IsMatch(text.Substring(1), pattern)));
            }
            else
            {
                return match && IsMatch(text.Substring(1), pattern.Substring(1));
            }
        }
    }
}

using System;
using System.Collections.Generic;
using System.Text;

namespace LeetCode.Problems.Medium.Problem49
{
    class Solution
    {
        public IList<IList<string>> GroupAnagrams(string[] strs)
        {
            if (strs == null || strs.Length == 0)
                return new List<IList<string>>();

            IList<IList<string>> anagrams = new List<IList<string>>();
            Dictionary<string, List<string>> dictionary = new Dictionary<string, List<string>>();

            foreach (var word in strs)
            {
                var letters = word.ToCharArray();
                Array.Sort(letters);

                if (!dictionary.ContainsKey(string.Join("", letters)))
                    dictionary.Add(string.Join("", letters), new List<string>());

                dictionary[string.Join("", letters)].Add(word);
            }

            foreach (var item in dictionary.Values)
                anagrams.Add(item);

            return anagrams;
        }
    }
}

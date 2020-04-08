using System;
using System.Collections.Generic;
using System.Linq;

namespace LeetCode.Problems._30DayChallenge.Counting_Elements
{
    class Solution
    {
        public int CountElements(int[] arr)
        {
            if (arr == null || arr.Length == 0)
                return 0;

            IDictionary<int, int> hashMap = new Dictionary<int, int>();

            for (int i = 0; i < arr.Length; i++)
            {
                if (!hashMap.ContainsKey(arr[i]))
                {
                    hashMap.Add(arr[i], 1);
                    continue;
                }

                hashMap[arr[i]] += 1;
            }

            var count = 0;
            var keyArray = hashMap.Keys.ToArray();
            Array.Sort(keyArray);

            for (int i = keyArray.Length - 1; 0 <= i; i--)
            {
                var key = keyArray[i];
                if (hashMap.ContainsKey(key + 1))
                {
                    for (int j = 0; j < hashMap[key]; j++)
                        count++;
                }
            }

            return count;
        }
    }
}

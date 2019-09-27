using System;
using System.Collections.Generic;

namespace LeetCode.Problems.Medium
{
    class Problem4
    {
        public double FindMedianSortedArrays(int[] nums1, int[] nums2)
        {
            var mergedArray = new List<int>();

            this.TryAddRange(mergedArray, nums1);
            this.TryAddRange(mergedArray, nums2);

            mergedArray.Sort();

            var medianIndex = mergedArray.Count / 2;

            if (mergedArray.Count % 2 == 1 || medianIndex == 0)
            {
                return mergedArray.Count > 0 ? mergedArray[medianIndex] : 0;
            }

            return (mergedArray[medianIndex - 1] + mergedArray[medianIndex]) / 2.0;
        }

        private void TryAddRange(List<int> array, int[] range)
        {
            if (range != null)
            {
                array.AddRange(range);
            }
        }
    }
}

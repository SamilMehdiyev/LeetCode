using System;
using System.Linq;

namespace LeetCode.Problems.Hard.Problem4
{
    class Solution3
    {
        public double FindMedianSortedArrays(int[] nums1, int[] nums2)
        {
            if (nums1 == null || (nums1 != null && nums1.Length == 0))
            {
                return this.processNotNullArray(nums2);
            }
            else if (nums2 == null || (nums2 != null && nums2.Length == 0))
            {
                return this.processNotNullArray(nums1);
            }

            if (nums1.Length == 1 && nums2.Length == 2)
            {
                return this.processSingleArray(nums1, nums2);
            }
            else if (nums2.Length == 1 && nums1.Length == 2)
            {
                return this.processSingleArray(nums2, nums1);
            }
            else if (nums1.Length == 1 && nums2.Length == 1)
            {
                return (nums1[0] + nums2[0]) / 2.0;
            }

            if (nums1.Length == 1 && nums2.Length > 2)
            {
                return this.processArray(nums1, nums2);
            }
            else if (nums2.Length == 1 && nums1.Length > 2)
            {
                return this.processArray(nums2, nums1);
            }

            int firstArrayLength = nums1.Length, secondArrayLength = nums2.Length,
                firstArrayCenter = firstArrayLength / 2, secondArrayCenter = secondArrayLength / 2;

            int totalLength = firstArrayLength + secondArrayLength;

            double firstArrayMedian = 0, secondArrayMedian = 0;

            while ((firstArrayLength >= 2) && (secondArrayLength >= 2))
            {
                firstArrayMedian = firstArrayLength % 2 == 0 ?
                    (nums1[firstArrayCenter - 1] + nums1[firstArrayCenter]) / 2.0 : nums1[firstArrayCenter];

                secondArrayMedian = secondArrayLength % 2 == 0 ?
                    (nums2[secondArrayCenter - 1] + nums2[secondArrayCenter]) / 2.0 : nums2[secondArrayCenter];

                if (firstArrayMedian == secondArrayMedian)
                {
                    return firstArrayMedian;
                }

                if (firstArrayMedian > secondArrayMedian)
                {
                    nums1 = nums1.Where(item => item <= firstArrayMedian).ToArray();
                    nums2 = nums2.Where(item => item >= secondArrayMedian).ToArray();
                }

                if (firstArrayMedian < secondArrayMedian)
                {
                    nums1 = nums1.Where(item => item >= firstArrayMedian).ToArray();
                    nums2 = nums2.Where(item => item <= secondArrayMedian).ToArray();
                }

                firstArrayLength = nums1.Length;
                secondArrayLength = nums2.Length;

                firstArrayCenter = firstArrayLength / 2;
                secondArrayCenter = secondArrayLength / 2;
            }

            if (nums1.Length == 1 && nums2.Length >= 2)
            {
                if (totalLength % 2 == 0)
                {
                    int median = nums2[0] + nums2[1];

                    for (int i = 0; i < nums2.Length; i++)
                    {
                        if (nums1[0] > nums2[i])
                        {
                            median = nums1[0] + nums2[i];
                        }
                    }

                    return median / 2.0;
                }
                else
                {
                    return processSingleArray(nums1, nums2);
                }
            }
            else if (nums2.Length == 1 && nums1.Length >= 2)
            {
                if (totalLength % 2 == 0)
                {
                    int median = nums1[0] + nums1[1];

                    for (int i = 0; i < nums1.Length; i++)
                    {
                        if (nums2[0] > nums1[i])
                        {
                            median = nums2[0] + nums1[i];
                        }
                    }

                    return median / 2.0;
                }
                else
                {
                    return processSingleArray(nums2, nums1);
                }
            }

            return (Math.Max(nums1[0], nums2[0]) + Math.Min(nums1[1], nums2[1])) / 2.0;
        }

        private double processNotNullArray(int[] array)
        {
            int length = array.Length;
            int center = length / 2;
            return length % 2 == 0 ? (array[center - 1] + array[center]) / 2.0 : array[center];
        }

        private double processSingleArray(int[] nums1, int[] nums2)
        {
            if (nums1[0] < nums2[0])
                return nums2[0];
            else if (nums1[0] > nums2[1])
                return nums2[1];
            else
                return nums1[0];
        }

        private double processArray(int[] firstArray, int[] secondArray)
        {
            int center = secondArray.Length / 2;
            int totalLength = secondArray.Length + 1;

            if (firstArray[0] <= secondArray[center] && firstArray[0] <= secondArray[center - 1])
            {
                return totalLength % 2 == 0 ? 
                    (secondArray[center - 1] + secondArray[center]) / 2.0 : secondArray[center - 1];
            }
            else if (firstArray[0] >= secondArray[center] && firstArray[0] >= secondArray[center + 1])
            {
                return totalLength % 2 == 0 ? 
                    (secondArray[center] + secondArray[center + 1]) / 2.0 : secondArray[center + 1];
            }
            else if (firstArray[0] <= secondArray[center] && firstArray[0] >= secondArray[center - 1])
            {
                return totalLength % 2 == 0 ? 
                    (firstArray[0] + secondArray[center]) / 2.0 : firstArray[0];
            }

            return totalLength % 2 == 0 ? (firstArray[0] + secondArray[center]) / 2.0 : secondArray[center];
        }
    }
}

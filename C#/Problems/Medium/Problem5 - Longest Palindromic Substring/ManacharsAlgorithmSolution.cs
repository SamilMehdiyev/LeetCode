using System;

namespace LeetCode.Problems.Medium.Problem5
{
    class ManacharsAlgorithmSolution
    {
        public string LongestPalindrome(string s)
        {
            int[] array = new int[(s.Length + 1) * 2];

            string newString = getNewString(s);

            int center = 0, right = 0;

            for (int index = 1; index < newString.Length - 1; index++)
            {
                int indexMirror = center - (index - center);

                if (right > index)
                {
                    array[index] = Math.Min(right - index, array[indexMirror]);
                }
                
                while (newString[index + 1 + array[index]] == newString[index - 1 - array[index]])
                {
                    array[index]++;
                }
                
                if (index + array[index] > right)
                {
                    center = index;
                    right = index + array[index];
                }
            }

            int maxPalindrome = 0;
            int centerIndex = 0;

            for (int i = 1; i < newString.Length - 1; i++)
            {

                if (array[i] > maxPalindrome)
                {
                    maxPalindrome = array[i];
                    centerIndex = i;
                }
            }

            return s.Substring((centerIndex - 1 - maxPalindrome) / 2, maxPalindrome);
        }

        private string getNewString(string s)
        {
            string newString = "@";

            for (int i = 0; i < s.Length; i++)
            {
                newString += "#" + s.Substring(i, 1);
            }

            newString += "#$";

            return newString;
        }
    }
}

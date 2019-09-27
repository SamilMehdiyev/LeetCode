

namespace LeetCode.Problems.Medium
{
    class Problem5
    {
        public string LongestPalindrome(string s)
        {
            var maxPalindromicSubString = "";

            if (s.Length == 0 || s.Length == 1)
                return s;

            var count = 1;
            var length = s.Length;

            while (true)
            {
                for (int i = 0; i < count; i++)
                {
                    var subStr = s.Substring(i, length);

                    if (isPandromicSubString(subStr))
                    {
                        maxPalindromicSubString = subStr;
                        break;
                    }
                }

                if (maxPalindromicSubString.Length > 0)
                {
                    break;
                }

                count++;
                length--;
            }

            return maxPalindromicSubString;
        }

        private bool isPandromicSubString(string str)
        {
            bool isPandomic = true;

            int rightInnerIndex = str.Length / 2;
            int leftInnerIndex = str.Length / 2;

            if (str.Length % 2 == 0)
            {
                rightInnerIndex = (str.Length / 2) - 1;
            }

            for (int leftSideIndex = 0; leftSideIndex < str.Length; leftSideIndex++)
            {
                int rightSideIndex = str.Length - (1 + leftSideIndex);

                if((leftSideIndex >= rightSideIndex) && isPandomic)
                {
                    return isPandomic;
                }

                if ((!str[leftSideIndex].Equals(str[rightSideIndex])) || (!str[leftInnerIndex].Equals(str[rightInnerIndex])))
                {
                    return false;
                }

                leftInnerIndex--;
                rightInnerIndex++;
            }

            return isPandomic;
        }
    }
}

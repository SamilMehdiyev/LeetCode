namespace LeetCode.Problems.Easy.Problem9
{
    class Solution
    {
        public bool IsPalindrome(int x)
        {
            if (x < 0)
            {
                return false;
            }
            int oldX = x;
            int palindromeX = 0;

            while(x > 0)
            {
                int digit = x % 10;

                x = x / 10;

                palindromeX = (10 * palindromeX) + digit;
            }

            if (oldX != palindromeX)
            {
                return false;
            }

            return true;
        }
    }
}

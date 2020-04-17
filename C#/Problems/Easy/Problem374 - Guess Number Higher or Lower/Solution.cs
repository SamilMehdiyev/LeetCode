namespace LeetCode.Problems.Easy.Problem374
{
    class Solution
    {
        public int GuessNumber(int n)
        {
            long start = 0;
            long end = n;
            long mid = 0;

            while (start <= end)
            {
                mid = (start + end) / 2;
                var result = guess((int)mid);

                if (result == -1)
                    end = mid - 1;
                else if (result == 1)
                    start = mid + 1;
                else
                    return (int)mid;
            }

            return (int)mid;
        }

        private int guess(int num)
        {
            var guess = 6;

            if (num < guess)
                return -1;
            else if (num > guess)
                return 1;
            else
                return 0;
        }
    }
}

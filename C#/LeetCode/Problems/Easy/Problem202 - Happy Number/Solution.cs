namespace LeetCode.Problems.Easy.Problem202
{
    class Solution
    {
        public bool IsHappy(int n)
        {
            var sum = 0;
            var result = false;
            var attempt = 10;

            while(attempt > 0)
            {
                var reminder = n % 10;
                n = n / 10;

                sum += (reminder * reminder);

                if (n == 0 && sum == 1)
                {
                    result = true;
                    break;
                } else if (n == 0 && sum != 1)
                {
                    n = sum;
                    sum = 0;
                    attempt--;
                }
            }

            return result;
        }
    }
}

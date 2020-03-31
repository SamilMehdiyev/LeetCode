namespace LeetCode.Problems.Easy.Problem278
{
    class Solution
    {
        public int FirstBadVersion(int n)
        {
            long start = 1;
            long end = n;

            while (start < end)
            {
                long mid = (start + end) / 2;

                var isBadVersion = IsBadVersion((int)mid);

                if (isBadVersion)
                    end = mid;
                else
                    start = mid + 1;
            }

            if (start != n && IsBadVersion((int)start)) return (int)start;
            return (int)end;
        }

        private bool IsBadVersion(int num)
        {
            var guess = 4;

            if (num < guess)
                return false;
            else if (num > guess)
                return true;
            else
                return true;
        }
    }
}

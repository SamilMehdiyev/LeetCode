namespace LeetCode.Problems.Easy.Problem69
{
    class Solution
    {
        public int MySqrt(int x)
        {
            var start = 0;
            var end = x;

            while (start <= end)
            {
                var mid = (long)(start + end) / 2;
                if (mid * mid == x)
                    return (int)mid;
                if (mid * mid < x)
                    start = (int)mid + 1;
                else
                    end = (int)mid - 1;
            }
            return start <= end ? start : end;
        }
    }
}

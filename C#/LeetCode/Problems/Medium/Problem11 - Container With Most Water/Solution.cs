namespace LeetCode.Problems.Medium.Problem11
{
    class Solution
    {
        public int MaxArea(int[] height)
        {
            var max = 0;

            for (int i = 1; i <= height.Length - 1; i++)
            {
                var j = i - 1;
                var x = 1;

                while (j >= 0)
                {
                    var y = height[i] < height[j] ? height[i] : height[j];

                    var result = y * x;

                    if (result > max)
                    {
                        max = result;
                    }

                    j--;
                    x++;
                }
            }
            return max;
        }
    }
}

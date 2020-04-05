namespace LeetCode.Problems.Easy.Problem121
{
    class Solution
    {
        public int MaxProfit(int[] prices)
        {
            if (prices == null || prices.Length < 2)
                return 0;

            var buyAt = 0;
            var sellAt = 0;
            var maxProfit = 0;

            for (int i = 1; i < prices.Length; i++)
            {
                var prev = prices[i - 1];
                var current = prices[i];

                if (prev > current)
                {
                    var profit = prices[sellAt] - prices[buyAt];

                    if (profit > maxProfit)
                    {
                        maxProfit = profit;
                        sellAt = 0;
                    }

                    if (prices[buyAt] > prices[i])
                        buyAt = i;
                }

                if (sellAt < buyAt)
                    sellAt = buyAt;

                sellAt = i;
            }
            
            return (prices[sellAt] - prices[buyAt]) > maxProfit ? (prices[sellAt] - prices[buyAt]) : maxProfit;
        }
    }
}

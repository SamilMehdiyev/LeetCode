namespace LeetCode.Problems.Easy.Problem122
{
    class Solution
    {
        public int MaxProfit(int[] prices)
        {
            if (prices == null || prices.Length < 2)
                return 0;

            var buyAt = -1;
            var sellAt = -1;
            var profit = 0;

            for (int i = 1; i < prices.Length; i++)
            {
                var prev = prices[i - 1];
                var current = prices[i];

                if (prev > current)
                {

                    if (sellAt != -1)
                    {
                        profit += prices[sellAt] - prices[buyAt];
                        buyAt = -1;
                        sellAt = -1;
                        continue;
                    }

                    buyAt = i;
                    continue;
                }

                sellAt = i;
                if (buyAt == -1)
                    buyAt = i - 1;
            }

            if (buyAt != -1 && sellAt != -1)
                return profit + (prices[sellAt] - prices[buyAt]);

            return profit;
        }
    }
}

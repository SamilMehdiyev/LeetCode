using System;
using System.Text;

namespace LeetCode.Problems.Easy.Problem7
{
    class Solution
    {
        public int Reverse(int x)
        {
            Console.WriteLine(x);

            StringBuilder reversedNumber = new StringBuilder("");
            bool isPositive = true;
            var index = 0;

            foreach (var item in x.ToString())
            {
                if (item.Equals('-') && isPositive)
                {
                    isPositive = false;
                    reversedNumber.Append(item);
                    index++;
                    continue;
                }

                reversedNumber.Insert(index, item);
            }

            var result = 0;

            int.TryParse(reversedNumber.ToString(), out result);

            return result;
        }
    }
}

using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy
{
    class Problem7
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

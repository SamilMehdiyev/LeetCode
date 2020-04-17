using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem13
{
    class Solution
    {
        public int RomanToInt(string s)
        {
            int prevNumber = this.getNumber(s[0]);
            int result = 0;

            result += prevNumber;

            for (int i = 1; i < s.Length; i++)
            {
                int number = this.getNumber(s[i]);

                if (prevNumber < number)
                {
                    result += (-1 * prevNumber) + (number - prevNumber);
                    prevNumber = number;

                    continue;
                }

                result += number;
                prevNumber = number;
            }

            return result < 0 ? result * -1 : result;
        }

        private int getNumber(char item)
        {
            switch (item)
            {
                case 'I': return 1;
                case 'V': return 5;
                case 'X': return 10;
                case 'L': return 50;
                case 'C': return 100;
                case 'D': return 500;
                case 'M': return 1000;
                default: return 0;
            }
        }
    }
}

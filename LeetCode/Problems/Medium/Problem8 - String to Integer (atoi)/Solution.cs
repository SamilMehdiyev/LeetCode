using System;

namespace LeetCode.Problems.Medium.Problem8
{
    class Solution
    {
        public int MyAtoi(string str)
        {
            if (string.IsNullOrWhiteSpace(str))
                return 0;

            int lastIndex = str.Length - 1;

            bool isPositive = true;
            bool lastElementIsNonDigit = false;

            string number = "";

            for (int i = 0; i < str.Length; i++)
            {
                char element = str[i];

                if (char.IsWhiteSpace(element))
                {
                    if (number.Length != 0)
                    {
                        return this.parseNumber(isPositive, number);
                    }

                    if (lastElementIsNonDigit)
                    {
                        return 0;
                    }

                    continue;
                }

                if (element == 45)
                {
                    if (!lastElementIsNonDigit && number.Length > 0)
                    {
                        return this.parseNumber(isPositive, number);
                    }

                    if (lastIndex != i)
                    {
                        isPositive = false;
                    }

                    if (str.Length == 1 || lastElementIsNonDigit)
                        return 0;

                    lastElementIsNonDigit = true;

                    continue;
                }

                if (element == 43)
                {
                    if (!lastElementIsNonDigit && number.Length > 0)
                    {
                        return this.parseNumber(isPositive, number);
                    }

                    if (lastIndex != i)
                    {
                        isPositive = true;
                    }

                    if (str.Length == 1 || lastElementIsNonDigit)
                        return 0;

                    lastElementIsNonDigit = true;

                    continue;
                }

                if (!char.IsDigit(element))
                {
                    if (number.Length != 0)
                    {
                        return this.parseNumber(isPositive, number);
                    }

                    return 0;
                }
                else
                {
                    if (lastElementIsNonDigit && number.Length != 0)
                    {
                        return this.parseNumber(isPositive, number);
                    }

                    number += element.ToString();

                    lastElementIsNonDigit = false;
                }
            }

            if (number.Length != 0)
            {
                return this.parseNumber(isPositive, number);
            }

            return 0;
        }



        private int parseNumber(bool isPositive, string number)
        {
            bool parsed = int.TryParse(number, out int result);

            if (parsed)
            {
                return isPositive ? result : -1 * result;
            }

            return isPositive ? int.MaxValue : int.MinValue;
        }
    }
}

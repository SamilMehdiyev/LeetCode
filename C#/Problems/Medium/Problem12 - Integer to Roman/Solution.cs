using System.Text;

namespace LeetCode.Problems.Medium.Problem12
{
    class Solution
    {
        private const int delimeter = 10;

        public string IntToRoman(int num)
        {
            if (num < 1 && num > 3999)
                return "";

            var result = new StringBuilder();
            var value = 1;

            while (num > 0)
            {
                var digit = (num % delimeter);

                var number = digit * value;

                switch (number)
                {
                    case 1:    result.Insert(0, "I"); break;
                    case 5:    result.Insert(0, "V"); break;
                    case 10:   result.Insert(0, "X"); break;
                    case 50:   result.Insert(0, "L"); break;
                    case 100:  result.Insert(0, "C"); break;
                    case 500:  result.Insert(0, "D"); break;
                    case 1000: result.Insert(0, "M"); break;
                }

                if (1 < number && number < 5)
                {
                    if (5 - number == 1)
                    {
                        result.Insert(0, "IV");
                    }
                    else
                    {
                        result.Insert(0, new string('I', digit));
                    }
                }

                if (5 < number && number < 10)
                {
                    if (10 - number == 1)
                    {
                        result.Insert(0, "IX");
                    }
                    else
                    {
                        result.Insert(0, "V" + new string('I', digit - 5));
                    }
                }

                if (10 < number && number < 50)
                {
                    if (50 - number == 10)
                    {
                        result.Insert(0, "XL");
                    }
                    else
                    {
                        result.Insert(0, new string('X', digit));
                    }
                }

                if (50 < number && number < 100)
                {
                    if (100 - number == 10)
                    {
                        result.Insert(0, "XC");
                    }
                    else
                    {
                        result.Insert(0, "L" + new string('X', digit - 5));
                    }
                }

                if (100 < number && number < 500)
                {
                    if (500 - number == 100)
                    {
                        result.Insert(0, "CD");
                    }
                    else
                    {
                        result.Insert(0, new string('C', digit));
                    }
                }

                if (500 < number && number < 1000)
                {
                    if (1000 - number == 100)
                    {
                        result.Insert(0, "CM");
                    }
                    else
                    {
                        result.Insert(0, "D" + new string('C', digit - 5));
                    }
                }

                if (number > 1000)
                {
                    result.Insert(0, new string('M', num));
                }


                value *= delimeter;
                num /= delimeter;

                continue;
            }

            return result.ToString();
        }
    }
}

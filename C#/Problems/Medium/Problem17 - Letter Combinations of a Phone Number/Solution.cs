using System.Collections.Generic;

namespace LeetCode.Problems.Medium.Problem17
{
    class Solution
    {
        private readonly IList<string> combinations = new List<string>();
        private string[][] dictionary;

        public IList<string> LetterCombinations(string digits)
        {
            if (digits.Length == 0)
            {
                return combinations;
            }
            
            if (digits.Length == 1)
            {
                foreach (var item in this.getStringArrayOfDigit(digits[0]))
                {
                    combinations.Add(item);
                }

                return combinations;
            }

            var index = 0;
            dictionary = new string[digits.Length][];

            foreach (var digit in digits)
            {
                dictionary[index] = this.getStringArrayOfDigit(digit);
                index++;
            }

            this.generateCombinations(0, "");

            return combinations;
        }


        private void generateCombinations(int index, string str)
        {
            foreach (var item in dictionary[index])
            {
                if (index < dictionary.Length - 1)
                {
                    this.generateCombinations(index + 1, str + item);
                }
                else
                {
                    combinations.Add(str + item);
                }
            }
        }

        private string[] getStringArrayOfDigit(char digit)
        {
            switch (digit)
            {
                case '2':
                    return new string[] { "a", "b", "c" };
                case '3':
                    return new string[] { "d", "f", "e" };
                case '4':
                    return new string[] { "g", "h", "i" };
                case '5':
                    return new string[] { "j", "k", "l" };
                case '6':
                    return new string[] { "m", "n", "o" };
                case '7':
                    return new string[] { "p", "q", "r", "s" };
                case '8':
                    return new string[] { "t", "u", "v" };
                case '9':
                    return new string[] { "w", "x", "y", "z" };
                default:
                    return new string[] { "" };
            }
        }
    }
}

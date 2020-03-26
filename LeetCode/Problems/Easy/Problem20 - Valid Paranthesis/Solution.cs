using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode.Problems.Easy.Problem20
{
    class Solution
    {
        public bool IsValid(string s)
        {
            if (s.Length == 0)
                return false;

            int[] array = changeStringToNumericArray(s);
            List<int> prevs = new List<int>();

            prevs.Add(array[0]);
            var prevIndex = 0;

            var sum = array[0];

            for (int i = 1; i < array.Length; i++)
            {
                var item = array[i];

                if (item < 0)
                {
                    if (prevs.Count == 0)
                        return false;

                    if (-1 * prevs[prevIndex] == item)
                    {
                        prevs.RemoveAt(prevIndex);
                        prevIndex--;

                        sum += item;
                        continue;
                    }
                    else
                    {
                        return false;
                    }

                }

                if (item > 0)
                {
                    sum += item;

                    prevIndex++;
                    prevs.Add(item);
                }
            }

            return sum == 0 ? true : false;
        }

        private int[] changeStringToNumericArray(string str)
        {
            int[] array = new int[str.Length];

            var index = 0;
            foreach (var item in str)
            {
                var number = 0;

                switch (item)
                {
                    case '(':
                        number = 1;
                        break;
                    case '{':
                        number = 2;
                        break;
                    case '[':
                        number = 3;
                        break;
                    case ')':
                        number = -1;
                        break;
                    case '}':
                        number = -2;
                        break;
                    case ']':
                        number = -3;
                        break;
                }

                array[index] = number;
                index++;
            }

            return array;
        }
    }
}


namespace LeetCode.Problems.Medium.Problem6
{
    class Solution
    {
        public string Convert(string s, int numRows)
        {
            char[,] matrix = new char[this.getMatriXAxisCount(s, numRows), numRows];

            matrix = this.fillMatrixWithValue(s, numRows);

            string result = "";

            for (int j = matrix.GetLength(1) - 1; j >= 0; j--)
            {
                for (int i = 0; i < matrix.GetLength(0); i++)
                    if (!matrix[i, j].Equals('\0'))
                    {
                        result += matrix[i, j];
                    }
            }
            return result;
        }

        private int getMatriXAxisCount(string str, int numRows)
        {
            var index = 0;
            var lastIndex = numRows - 1;

            var count = 0;

            while (index <= str.Length)
            {
                index += numRows;
                count++;

                for (int x = lastIndex - 1; x > 0 && index <= str.Length; x--)
                {
                    index++;
                    count++;
                }
            }

            return count;
        }

        private char[,] fillMatrixWithValue(string s, int numRows)
        {
            var index = 0;
            var lastIndex = numRows - 1;

            var column = 0;

            char[,] matrix = new char[this.getMatriXAxisCount(s, numRows), numRows];

            while (index < s.Length)
            {
                for (int y = lastIndex; y >= 0 && index < s.Length; y--)
                {
                    matrix[column, y] = s[index];
                    index++;
                }

                column++;

                for (int y = 1; y <= lastIndex - 1 && index < s.Length; y++)
                {
                    matrix[column, y] = s[index];
                    column++;
                    index++;
                }
            }

            return matrix;
        }
    }
}

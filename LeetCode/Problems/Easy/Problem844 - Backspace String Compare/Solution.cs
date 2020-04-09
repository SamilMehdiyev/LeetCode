namespace LeetCode.Problems.Easy.Problem844
{
    class Solution
    {
        public bool BackspaceCompare(string S, string T)
        {
            S = this.processString(S);
            T = this.processString(T);
            return S.Equals(T);
        }

        private string processString(string str)
        {
            var result = "";

            for (int i = 0; i < str.Length - 1; i++)
            {
                if (str[i].Equals('#'))
                {
                    if (!result.Equals(""))
                        result = result.Substring(0, result.Length - 1);

                    continue;
                }

                if (i + 1 < str.Length && str[i + 1].Equals('#'))
                {
                    i += 1;
                    continue;
                }

                result += str[i];
            }

            return result;
        }
    }
}

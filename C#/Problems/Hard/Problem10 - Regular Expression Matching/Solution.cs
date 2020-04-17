using System.Linq;

// Not finished
namespace LeetCode.Problems.Hard.Problem10
{
    class Solution
    {
        public bool IsMatch(string s, string p)
        {
            if (p.Equals(".*"))
                return true;

            if (p.Equals("*"))
                return false;

            if (s.Length > p.Length && p.Length == 1)
                return false;

            if (string.IsNullOrWhiteSpace(s) || string.IsNullOrWhiteSpace(p))
                return false;

            int sIndex = 0;
            int pIndex = 0;

            while (sIndex < s.Length && pIndex < p.Length)
            {
                var sElement = s[sIndex];
                var pElement = p[pIndex];
                var pNextElement = ((pIndex + 1) < p.Length) ? p[pIndex + 1] : '$';

                if (sElement.Equals(pElement) && pNextElement.Equals('$'))
                {
                    sIndex++;
                    pIndex++;
                    break;
                }

                if (pElement.Equals('.') && pNextElement.Equals('*'))
                {
                    int sameElementCount = 0;

                    for (int i = sIndex; i < s.Length; i++)
                    {
                        if (!s[i].Equals(sElement))
                        {
                            break;
                        }

                        sameElementCount++;
                    }

                    if (sameElementCount > 2)
                    {
                        pIndex += 2;

                        for (int i = 0; i <= sameElementCount; i++)
                        {
                            string replace = new string(sElement, i);

                            string pNew = p.Substring(0, pIndex - 2) + replace + p.Substring(pIndex);

                            bool result = IsMatch(s, pNew);

                            if (result)
                            {
                                return true;
                            }
                        }
                    }
                    else
                    {
                        sIndex += (s.Length - sIndex);
                        pIndex += 2;
                        break;
                    }
                }

                if (pElement.Equals('.'))
                {
                    sIndex++;
                    pIndex++;
                    continue;
                }

                if (!sElement.Equals(pElement) && !pNextElement.Equals('*'))
                    return false;

                if (!sElement.Equals(pElement) && pNextElement.Equals('*'))
                {
                    pIndex += 2;
                    continue;
                }

                if (sElement.Equals(pElement) && !pNextElement.Equals('*'))
                {
                    sIndex++;
                    pIndex++;
                    continue;
                }

                if (sElement.Equals(pElement) && pNextElement.Equals('*'))
                {
                    int sameElementCount = 0;

                    for (int i = sIndex; i < s.Length; i++)
                    {
                        if (!s[i].Equals(sElement))
                        {
                            break;
                        }

                        sameElementCount++;
                    }
                
                    pIndex += 2;

                    for (int i = 0; i <= sameElementCount; i++)
                    {
                        string replace = new string(sElement, i);

                        string pNew = p.Substring(0, pIndex - 2) + replace + p.Substring(pIndex);

                        bool result = IsMatch(s, pNew);

                        if (result)
                        {
                            return true;
                        }
                    }
                }
            }

            if (s.Length == sIndex && p.Length == pIndex)
                return true;

            if (p.Substring(pIndex).Contains('*') && !p.Contains('.'))
            {
                return IsMatch(s, p.Substring(pIndex));
            }

            return false;
        }
    }
}

using System;
using System.Numerics;

namespace LeetCode
{
    class Program
    {
        static void Main(string[] args)
        {
            processProblem11();
        }

        public static void processProblem2()
        {
            Problems.Medium.Problem2.Solution problem = new Problems.Medium.Problem2.Solution();
            Problems.Medium.Problem2.ListNode result = problem.AddTwoNumbers(BigInteger.Parse("1000000000000000000000000000001"), BigInteger.Parse("564"));
            Console.WriteLine(problem.getValue(result));
        }

        public static void processProblem3()
        {
            Problems.Medium.Problem3.Solution problem = new Problems.Medium.Problem3.Solution();
            Console.WriteLine(problem.LengthOfLongestSubstring(" "));
        }

        public static void processProblem4()
        {
            //Problems.Hard.Problem4.Solution problem = new Problems.Hard.Problem4.Solution();
            //Problems.Hard.Problem4.Solution2 problem = new Problems.Hard.Problem4.Solution2();
            Problems.Hard.Problem4.Solution3 problem = new Problems.Hard.Problem4.Solution3();
            Console.WriteLine(problem.FindMedianSortedArrays(new[] { 1, 2 }, new[] { 3, 4, 5, 6, 7, 8, 9 }));
        }

        public static void processProblem5()
        {
            Problems.Medium.Problem5.ManacharsAlgorithmSolution problem = new Problems.Medium.Problem5.ManacharsAlgorithmSolution();
            Console.WriteLine(problem.LongestPalindrome("gphyvqruxjmwhonjjrgumxjhfyupajxbjgthzdvrdqmdouuukeaxhjumkmmhdglqrrohydrmbvtuwstgkoby" +
                "zjjtdtjroqpyusfsbjlusekghtfbdctvgmqzeybn" +
                "wzlhdnhwzptgkzmujfldoiejmvxnorvbiubfflygrkedyirienybosqzrkbpcfidvkkafftgzwrcitqizelhfsruwmtrgaocjcyxdkovtdennrkmxwpdsxpxuarh" +
                "gusizmwakrmhdwcgvfljhzcskclgrvvbrkesojyhofwqiwhiupujmkcvlywjtmbncurxxmpdskupyvvweuhbsnanzfioirecfxvmgcpwrpmbhmkdtckhvbxnsbci" +
                "fhqwjjczfokovpqyjmbywtpaqcfjowxnmtirdsfeujyogbzjnjcmqyzciwjqxxgrxblvqbutqittroqadqlsdzihngpfpjovbkpeveidjpfjktavvwurqrgqdomi" +
                "ibfgqxwybcyovysydxyyymmiuwovnevzsjisdwgkcbsookbarezbhnwyqthcvzyodbcwjptvigcphawzxouixhbpezzirbhvomqhxkfdbokblqmrhhioyqubpyqh" +
                "jrnwhjxsrodtblqxkhezubprqftrqcyrzwywqrgockioqdmzuqjkpmsyohtlcnesbgzqhkalwixfcgyeqdzhnnlzawrdgskurcxfbekbspupbduxqxjeczpmdvss" +
                "ikbivjhinaopbabrmvscthvoqqbkgekcgyrelxkwoawpbrcbszelnxlyikbulgmlwyffurimlfxurjsbzgddxbgqpcdsuutfiivjbyqzhprdqhahpgenjkbiukur" +
                "vdwapuewrbehczrtswubthodv"));
        }

        public static void processProblem6()
        {
            Problems.Medium.Problem6.Solution problem = new Problems.Medium.Problem6.Solution();
            problem.Convert("PAYPALISHIRING", 4);
        }

        public static void processProblem7()
        {
            Problems.Easy.Problem7.Solution problem = new Problems.Easy.Problem7.Solution();
            Console.WriteLine(problem.Reverse(120));
        }

        public static void processProblem8()
        {
            Problems.Medium.Problem8.Solution solution = new Problems.Medium.Problem8.Solution();
            Console.WriteLine(solution.MyAtoi("0-1"));
        }

        public static void processProblem9()
        {
            Problems.Easy.Problem9.Solution solution = new Problems.Easy.Problem9.Solution();
            Console.WriteLine(solution.IsPalindrome(101));
        }

        public static void processProblem10()
        {
            Problems.Hard.Problem10.Solution2 solution = new Problems.Hard.Problem10.Solution2();
            Console.WriteLine(solution.IsMatch("mississippi", "mis*is*p*."));
        }

        public static void processProblem11()
        {
            Problems.Medium.Problem11.Solution solution = new Problems.Medium.Problem11.Solution();
            Console.WriteLine(solution.MaxArea(new int[] { 1, 8, 6, 2, 5, 4, 8, 3, 7 }));
        }
    }
}

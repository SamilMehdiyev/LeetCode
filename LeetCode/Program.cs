using System;
using System.Numerics;

namespace LeetCode
{
    class Program
    {
        static void Main(string[] args)
        {
            processProblem2();
        }

        public static void processProblem2()
        {
            Problem2 q2 = new Problem2();
            ListNode result = q2.AddTwoNumbers(BigInteger.Parse("1000000000000000000000000000001"), BigInteger.Parse("564"));
            Console.WriteLine(q2.getValue(result));
        }
    }
}

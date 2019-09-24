using System;
using System.Collections.Generic;
using System.Linq;
using System.Numerics;
using System.Text;
using System.Threading.Tasks;

namespace LeetCode
{
    class Program
    {
        static void Main(string[] args)
        {
            Question2 q2 = new Question2();
            ListNode result = q2.AddTwoNumbers(BigInteger.Parse("1000000000000000000000000000001"), BigInteger.Parse("564"));
            Console.WriteLine(q2.getValue(result));
        }
    }
}

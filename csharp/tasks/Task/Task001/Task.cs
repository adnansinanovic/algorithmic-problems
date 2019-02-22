using System;

namespace tasks.Task.Task001
{
    public class Task001 : ICodeTask
    {
        public string Title => "Binary Gap";

        public string ShortDescription => "Find longest sequence of zeros in binary representation of an integer.\n";

        public string Description => "A binary gap within a positive integer N is any maximal sequence of consecutive zeros that is surrounded by ones at both ends in the binary representation of N.\n\nFor example, number 9 has binary representation 1001 and contains a binary gap of length 2. The number 529 has binary representation 1000010001 and contains two binary gaps: one of length 4 and one of length 3. The number 20 has binary representation 10100 and contains one binary gap of length 1. The number 15 has binary representation 1111 and has no binary gaps. The number 32 has binary representation 100000 and has no binary gaps.\n\nWrite a function:\n\nclass Solution { public int solution(int N); }\n\nthat, given a positive integer N, returns the length of its longest binary gap. The function should return 0 if N doesn't contain a binary gap.\n\nFor example, given N = 1041 the function should return 5, because N has binary representation 10000010001 and so its longest binary gap is of length 5. Given N = 32 the function should return 0, because N has binary representation '100000' and thus no binary gaps.\n\nWrite an efficient algorithm for the following assumptions:\n\nN is an integer within the range [1..2,147,483,647].\n";

        public string URL => "https://app.codility.com/programmers/lessons/1-iterations/binary_gap/";

        public void Show()
        {
            test(529, 4);
            test(1000, 1);
            test(20, 1);
            test(1041, 5);
            test(1041, 5);
            test(32, 0);
        }

        void test(int n, int expected)
        {
            var s = new Solution();
            int actual = s.solution(n);

            var binary = s.toBinary(n);
            if (actual != expected)
            {
                Console.WriteLine("Error: {0,5} :: {1,15} :: {2,10} <> {3} !!!!!!!!!!!!!!!!!!!!!", n, binary, actual, expected);
            }
            else
            {
                Console.WriteLine("OK: {0,5} :: {1,15} :: {2,10} == {3}", n, binary, actual, expected);
            }
        }
    }
}

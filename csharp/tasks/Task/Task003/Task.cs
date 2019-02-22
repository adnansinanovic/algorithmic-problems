using System;
using System.Text;

namespace tasks.Task.Task003
{
    public class Task003 : ICodeTask
    {
        public string Title => "Cyclic Rotation";

        public string ShortDescription => "Rotate an array to the right by a given number of steps.\n";

        public string Description => "An array A consisting of N integers is given. Rotation of the array means that each element is shifted right by one index, and the last element of the array is moved to the first place. For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7] (elements are shifted right by one index and 6 is moved to the first place).\r\n\r\nThe goal is to rotate array A K times; that is, each element of A will be shifted to the right K times.\r\n\r\nWrite a function:\r\n\r\nclass Solution { public int[] solution(int[] A, int K); }\r\n\r\nthat, given an array A consisting of N integers and an integer K, returns the array A rotated K times.\r\n\r\nFor example, given\r\n\r\n    A = [3, 8, 9, 7, 6]\r\n    K = 3\r\nthe function should return [9, 7, 6, 3, 8]. Three rotations were made:\r\n\r\n    [3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]\r\n    [6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]\r\n    [7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]\r\nFor another example, given\r\n\r\n    A = [0, 0, 0]\r\n    K = 1\r\nthe function should return [0, 0, 0]\r\n\r\nGiven\r\n\r\n    A = [1, 2, 3, 4]\r\n    K = 4\r\nthe function should return [1, 2, 3, 4]\r\n\r\nAssume that:\r\n\r\nN and K are integers within the range [0..100];\r\neach element of array A is an integer within the range [−1,000..1,000].\r\nIn your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.";

        public string URL => "https://app.codility.com/programmers/lessons/2-arrays/cyclic_rotation/";

        public void Show()
        {
            Test(new int[] { 1, 2, 3, 4, 5 }, 2, new int[] { 4, 5, 1, 2, 3 });
            Test(new int[] { 3, 8, 9, 7, 6 }, 3, new int[] { 9, 7, 6, 3, 8 });
            Test(new int[] {  }, 3, new int[] {  });
            //[3, 8, 9, 7, 6]
        }

        private void Test(int[] original, int k, int[] expected)
        {
            var error = false;
            var s = new Solution();
            var actual = s.solution(original, k);

            string stringActual = string.Join(",", actual);
            string stringExpected = string.Join(",", expected);

            StringBuilder sb = new StringBuilder();
            if (actual.Length != expected.Length)
            {
                error = true;
                sb.AppendLine($"Lenght mismatch: {original.Length} <> {expected.Length}");
            }

            for (int i = 0; i < original.Length; i++)
            {
                var v1 = original[i];
                var v2 = expected[i];

                if (v1 != v2)
                {
                    error = true;
                    sb.AppendLine($"Value mismatch at index={i}: {v1} <> {v2}");
                }
            }

            Console.WriteLine($"Actual: {stringActual}");
            Console.WriteLine($"Expected: {stringExpected}");
            if (!error)
            {
                Console.WriteLine("Everything ok ");
            }
            else
            {
                Console.WriteLine(sb.ToString());
            }

            Console.WriteLine("==================================");
        }
    }
}
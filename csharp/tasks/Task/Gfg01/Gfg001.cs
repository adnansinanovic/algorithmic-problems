

using System;
using System.Linq;

namespace tasks.Task.Gfg001
{
    public class Gfg001 : ICodeTask
    {
        public string ShortDescription => "Find a sorted subsequence of size 3 in linear time";

        public string Title => "Sorted subsequence of size=3 in linear time";

        public string Description => "Given an array of n integers, find the 3 elements such that a[i] < a[j] < a[k] and i < j < k in 0(n) time. If there are multiple such triplets, then print any one of them.";

        public string URL => "https://www.geeksforgeeks.org/find-a-sorted-subsequence-of-size-3-in-linear-time/#comment-1573367479";

        public void Show()
        {
            Test(new int[] { 12, 11, 10, 5, 6, 2, 30 }, new int[] { 5, 6, 30 });
            Test(new int[] { 4, 3, 2, 1 }, new int[] { });
            Test(new int[] { 1, 2, 3, 4 }, new int[] { 1, 2, 3 });
        }

        private void Test(int[] arr, int[] expected)
        {
            var actual = Solution.solution2(arr);

            if (actual.Length == 0)
            {
                Console.WriteLine($"No such triplet found: expected={string.Join(",", expected)}, arr={string.Join(",", arr)} ");
            }
            else
            {
                Console.WriteLine($"ok actual={string.Join(",", actual)}, expected={string.Join(",", expected)}, arr={string.Join(",", arr)} ");
            }

            Console.WriteLine("=".PadRight(30, '='));

        }
    }
}



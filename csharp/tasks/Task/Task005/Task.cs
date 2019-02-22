using System;
namespace tasks.Task.Task005
{
    public class Task005 : ICodeTask
    {
        public string ShortDescription => "Find the missing element in a given permutation.";

        public string Title => "PermMissingElem";

        public string Description => "An array A consisting of N different integers is given. " +
        	"The array contains integers in the range [1..(N + 1)], which means that exactly one element is missing.\n\n" +
        	"Your goal is to find that missing element.\n\n" +
        	"Write a function:\n\n" +
        	"int solution(int A[], int N);\n\n" +
        	"that, given an array A, " +
        	"returns the value of the missing element.\n\n" +
        	"For example, given array A such that:\n\n  " +
        	"A[0] = 2\n  A[1] = 3\n  A[2] = 1\n  A[3] = 5\n" +
        	"the function should return 4, as it is the missing element.\n\n" +
        	"Write an efficient algorithm for the following assumptions:\n\n" +
        	"N is an integer within the range [0..100,000];\n" +
        	"the elements of A are all distinct;\n" +
        	"each element of array A is an integer within the range [1..(N + 1)";

        public string URL => "https://app.codility.com/programmers/lessons/3-time_complexity/perm_missing_elem/";

        public void Show()
        {
            Test(new int[] { 1, 2, 3, 4 }, 5);
            Test(new int[] { 1, 2, 3, 5 }, 4);
            Test(new int[] { 1, 2, 3, 4, 5, 6, 8, 9 }, 6);
        }

        private void Test(int[] n, int expceted)
        {
            var s = new Solution();
            var actual = s.solution(n);

            var arr = string.Join(", ", n);
            Console.WriteLine($"{(actual == expceted ? "OK" : "Error")}, actual={actual}, expected={expceted}, array={arr}");
        }
    }
}

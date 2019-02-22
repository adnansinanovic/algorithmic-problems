using System;

namespace tasks.Task.Task002
{
    class Task002 : ICodeTask
    {
        public string Title => "OddOccurancesInarray";

        public string ShortDescription => "Find value that occurs in odd number of elements.\n";

        public string Description => "A non-empty array A consisting of N integers is given. The array contains an odd number of elements, \n and each element of the array can be paired with another element that has the same value, \n except for one element that is left unpaired.\n\nFor example, in array A such that:\n\n  A[0] = 9  A[1] = 3  A[2] = 9\n  A[3] = 3  A[4] = 9  A[5] = 7\n  A[6] = 9\nthe elements at indexes 0 and 2 have value 9,\nthe elements at indexes 1 and 3 have value 3,\nthe elements at indexes 4 and 6 have value 9,\nthe element at index 5 has value 7 and is unpaired.\nWrite a function:\n\nint solution(int A[], int N);\n\nthat, given an array A consisting of N integers fulfilling the above conditions, returns the value of the unpaired element.\n\nFor example, given array A such that:\n\n  A[0] = 9  A[1] = 3  A[2] = 9\n  A[3] = 3  A[4] = 9  A[5] = 7\n  A[6] = 9\nthe function should return 7, as explained in the example above.\n\nWrite an efficient algorithm for the following assumptions:\n\nN is an odd integer within the range [1..1,000,000];\neach element of array A is an integer within the range [1..1,000,000,000];\nall but one of the values in A occur an even number of times.";

        public string URL => "https://app.codility.com/programmers/lessons/2-arrays/odd_occurrences_in_array/";

        public void Show()
        {
            Test(5, new int[] { 3, 2, 5, 2, 3 });
        }


        private void Test(int expected, int[] n)
        {
            var s = new Solution();

            int actual = s.solution(n);

            string array = string.Join(",", n);

            if (actual == expected)
            {
                Console.WriteLine("OK: {0,10} :: {1}", expected, array);
            }
            else
            {
                Console.WriteLine("ERROR: {0,10} <> {1} :: int[]{{{2}}} !!!!!!!!!!!", actual, expected, array);
            }
        }
    }
}

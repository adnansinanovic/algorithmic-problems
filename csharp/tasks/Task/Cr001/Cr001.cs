using System;
namespace tasks.Task.Cr001
{
    public class Cr001 : ICodeTask
    {
        public string ShortDescription => "";

        public string Title =>"Sock Merchant";

        public string Description => "John works at a clothing store. He has a large pile of socks that he must pair by color for sale. Given an array of integers representing the color of each sock, determine how many pairs of socks with matching colors there are.\n\nFor example, there are  socks with colors . There is one pair of color  and one of color . There are three odd socks left, one of each color. The number of pairs is .\n\nFunction Description\n\nComplete the sockMerchant function in the editor below. It must return an integer representing the number of matching pairs of socks that are available.\n\nsockMerchant has the following parameter(s):\n\nn: the number of socks in the pile\nar: the colors of each sock\nInput Format\n\nThe first line contains an integer , the number of socks represented in . \nThe second line contains  space-separated integers describing the colors  of the socks in the pile.\n\nConstraints\n\n where \nOutput Format\n\nReturn the total number of matching pairs of socks that John can sell.\n\nSample Input\n\n9\n10 20 20 10 10 30 50 10 20\nSample Output\n\n3\nExplanation\n\nsock.png\n\nJohn can match three pairs of socks.";

        public string URL => "https://www.hackerrank.com/challenges/sock-merchant/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=warmup";

        public void Show()
        {
            Test(new int[] { 10, 10, 20, 20, 30 }, 2);
            Test(new int[] { 10, 20, 20, 10, 10, 30, 50, 10, 20 }, 3);


        }

        private void Test(int[] n, int expected)
        {
            var actual  = Solution.sockMerchant(n.Length, n);

            var arr = string.Join(",", n);
            Console.WriteLine($"{(actual == expected ? "OK" : "Error")} actual={actual}, expected={expected}, arr={arr} ");
        }
    }
}

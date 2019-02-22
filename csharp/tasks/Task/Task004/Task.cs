using System;
namespace tasks.Task.Task004
{
    public class Task004 : ICodeTask
    {
        public string ShortDescription => "Count minimal number of jumps from position X to Y.\n";

        public string Title => "FrogJmp";

        public string Description => "A small frog wants to get to the other side of the road. The frog is currently located at position X and wants to get to a position greater than or equal to Y. The small frog always jumps a fixed distance, D.\n\nCount the minimal number of jumps that the small frog must perform to reach its target.\n\nWrite a function:\n\nclass Solution { public int solution(int X, int Y, int D); }\n\nthat, given three integers X, Y and D, returns the minimal number of jumps from position X to a position equal to or greater than Y.\n\nFor example, given:\n\n  X = 10\n  Y = 85\n  D = 30\nthe function should return 3, because the frog will be positioned as follows:\n\nafter the first jump, at position 10 + 30 = 40\nafter the second jump, at position 10 + 30 + 30 = 70\nafter the third jump, at position 10 + 30 + 30 + 30 = 100\nWrite an efficient algorithm for the following assumptions:\n\nX, Y and D are integers within the range [1..1,000,000,000];\nX ≤ Y.\n";

        public string URL => "https://app.codility.com/programmers/lessons/3-time_complexity/frog_jmp/";

        public void Show()
        {
            Test(10, 80, 30, 3);
            Test(10, 10, 30, 0);
            Test(2, 11, 3, 3);
        }

        private void Test(int x, int y, int d, int expected)
        {
            var s = new Solution();
            var actual = s.solution(x, y, d);

            Console.WriteLine($"{(actual == expected ? "OK" : "ERROR")} x={x}, y={y}, d={d}, actual={actual}, expected={expected}");
        }
    }
}

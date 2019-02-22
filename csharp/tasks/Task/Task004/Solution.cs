using System;
namespace tasks.Task.Task004
{
    public class Solution
    {
        public int solution(int X, int Y, int D)
        {
            int distance = Y - X;

            decimal steps = (decimal)distance / D;
            if (distance % D == 0)
            {
                return (int)steps;
            }
            else
            {
                return (int)Math.Round(steps + 0.5m);
            }
        }
    }
}

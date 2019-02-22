using System;
namespace tasks.Task.Task003
{
    public class Solution
    {
        public int[] solution(int[] A, int K)
        {
            if (K == 0 || A.Length == 0)
                return A;

            int[] res = new int[A.Length];
            int len = A.Length;
            for (int i = 0; i < len; i++)
            {
                if (i < (K % len))
                    res[i] = A[(len - (K % len)) + i];
                else
                    res[i] = A[System.Math.Abs((K % len) - i)];
            }

            return res;
            //var aLen = A.Length;
            //if (aLen == 0 || K == 0) {
            //    return A;
            //}

            //// store 
            //int[] swap = new int[K];
            //for (int i = 0; i < K; i++)
            //{
            //    swap[i] = A[aLen - K + i];
            //}

            //for (int i = aLen - 1; i >= K; i--)
            //{
            //    A[i] = A[i - K];
            //}


            //for (int i = 0; i < swap.Length; i++)
            //{
            //    A[i] = swap[i];
            //}

            //return A;
        }
    }
}

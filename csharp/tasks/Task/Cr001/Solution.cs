using System;
using System.Collections.Generic;
using System.IO;

namespace tasks.Task.Cr001
{
    class Solution
    {

        // Complete the sockMerchant function below
        public static int sockMerchant(int n, int[] ar)
        {
            Dictionary<int, int> socks = new Dictionary<int, int>();
            var len = ar.Length;
            for (int i = 0; i < len; i++)
            {
                var key = ar[i];
                if (socks.ContainsKey(key))
                {
                    socks[key]++;
                }
                else
                {
                    socks.Add(key, 1);
                }

            }

            var count = 0;
            foreach (var item in socks)
            {
                count += item.Value / 2;
            }

            return count;
        }

    }
}

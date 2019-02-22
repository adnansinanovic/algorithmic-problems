using System;
namespace tasks.Task.Task002
{
    class Solution
    {
        public int solution(int[] N)
        {
            int unMatchNumber = N[0];
            for (int i = 1; i < N.Length; i++)
            {
                unMatchNumber = unMatchNumber ^ N[i];
            }
            return unMatchNumber;
            //Dictionary<int, int> d = new Dictionary<int, int>();
            //for (int i = 0; i < N.Length; i++)
            //{
            //    var v = N[i];
            //    if (d.ContainsKey(v))
            //    {
            //        d[v]++;
            //    } else
            //    {
            //        d[v] = 1;f
            //    }
            //}


            //foreach (var item in d)
            //{
            //    if (item.Value % 2 == 1) {
            //        return item.Key;
            //    }
            //}

            //return 0;
        }
    }
}

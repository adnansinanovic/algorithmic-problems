using System;
namespace tasks.Task.Task001
{
    class Solution
    {
        public string toBinary(int n)
        {
            string s = "";
            while (n > 0)
            {
                var r = (n % 2);
                if (!(s.Length == 0 && r == 0))
                {
                    s += r;
                }
                n = n / 2;
            }

            return s;
        }

        public int solution(int N)
        {
            //var bin = Convert.ToString(N, 2);
            var bin = toBinary(N);
            int max = 0;
            int counter = 0;
            int len = bin.Length;
            for (int i = 0; i < len; i++)
            {
                char c = bin[i];
                if (c == '0')
                {
                    counter++;
                }
                else
                {
                    if (counter > max)
                    {
                        max = counter;
                    }

                    counter = 0;
                }
            }

            return max;

        }
    }
}

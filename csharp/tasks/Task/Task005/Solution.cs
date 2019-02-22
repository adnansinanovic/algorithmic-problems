using System;
using System.Collections.Generic;

namespace tasks.Task.Task005
{
    public class Solution
    {
        public int solution(int[] A)
        {
            throw new Exception("Not done");
            List<int> list = new List<int>(A);
            list.Sort();

            bool found = false;
            var startIndex = 0;
            var endIndex = list.Count - 1;
            while (found == false)
            {
                var count = endIndex - startIndex + 1;
                if (count <= 3) 
                {
                    for (int i = startIndex; i <= endIndex; i++)
                    {
                        if (list[startIndex] != startIndex - 1)
                        {
                            return startIndex - 1;
                        }
                    }

                    return -1;
                }
                else
                {
                    var middleIndex = startIndex + count / 2;
                    if (middleIndex + 1 == list[middleIndex])
                    {
                        startIndex = middleIndex + 1;
                    }
                    else
                    {
                        endIndex = middleIndex - 1;
                    }
                }

            }

            return -2;
        }
    }
}

using System;
using System.Collections.Generic;
using System.IO;

namespace tasks.Task.Gfg001
{
    class Solution
    {
        // Complete the sockMerchant function below
        public static int[] solution(int[] arr)
        {
            int n = arr.Length;

            // Index of maximum element from right side 
            int max = n - 1;

            // Index of minimum element from left side 
            int min = 0;
            int i;

            // Create an array that will store index  
            // of a smaller element on left side.  
            // If there is no smaller element 
            // on left side, then smaller[i] will be -1. 
            int[] smaller = new int[n];

            // first entry will always be -1 
            smaller[0] = -1;
            for (i = 1; i < n; i++)
            {
                if (arr[i] <= arr[min])
                {
                    min = i;
                    smaller[i] = -1;
                }
                else
                    smaller[i] = min;
            }

            // Create another array that will store  
            // index of a greater element on right side. 
            // If there is no greater element on  
            // right side, then greater[i] will be -1. 
            int[] greater = new int[n];

            // last entry will always be -1 
            greater[n - 1] = -1;
            for (i = n - 2; i >= 0; i--)
            {
                if (arr[i] >= arr[max])
                {
                    max = i;
                    greater[i] = -1;
                }
                else
                    greater[i] = max;
            }

            Console.WriteLine($"Smaller; {string.Join(",", smaller)}");
            Console.WriteLine($"Greater; {string.Join(",", greater)}");

            // Now find a number which has both a greater number 
            // on right side and smaller number on left side 
            for (i = 0; i < n; i++)
            {
                if (smaller[i] != -1 && greater[i] != -1)
                {
                    return new int[] { arr[smaller[i]], arr[i], arr[greater[i]] };
                }
            }

            // If we reach number, then there are no such 3 numbers 

            return new int[] { };
        }

        internal static object solution2(int[] arr)
        {

        }
    }
}

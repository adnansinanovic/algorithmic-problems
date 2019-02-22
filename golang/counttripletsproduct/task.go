//

package counttripletsproduct

import (
	"fmt"
	"strings"
)

func Do() {
	test([]int{2, 3, 6, 1, 6, 4, 4, 12, 24}, 12)

	// test([]int{1, 4, 6, 2, 3, 8}, 24, 3)
	// test([]int{0, 4, 6, 2, 3, 8}, 18, 0)

}

func test(arr []int, expected int) {
	actual := count2(arr)

	err := ""
	if actual != expected {
		err = strings.Repeat("!", 20)
	}

	fmt.Printf("arr=%v act=%v expe=%v %v\n",
		arr, actual, expected, err)
}

func count2(a []int) int {
	ln := len(a) - 1
	count := 0

	// 2, 3, 6, 1, 6, 4, 4, 12, 24
	for ln > 2 {
		doubleCandidates := map[int]int{}
		tripleCandidates := map[int]int{}

		for i := 0; i <= ln; i++ {
			v := a[i]
			count += tripleCandidates[v]

			for key, vv := range doubleCandidates {
				if key%v == 0 {
					tripleCandidates[key/v] += vv
				}
			}

			product := a[ln]
			if _, ok := doubleCandidates[product/v]; ok {
				doubleCandidates[product/v] = doubleCandidates[product/v] * 2
			} else {
				if product%v == 0 {
					doubleCandidates[product/v] = 1
				}
			}
		}
		ln--
	}
	return count
}

func count(a []int) int {
	// brute force
	var count int
	ln := len(a)

	for p := ln - 1; p > 3; p-- {
		product := a[p]
		for i := 0; i < ln; i++ {
			for ii := i + 1; ii < ln; ii++ {
				for iii := ii + 1; iii < ln; iii++ {
					if product == a[i]*a[ii]*a[iii] {
						fmt.Printf("%v x %v x %v = %v :: indices=%v,%v,%v\n", a[i], a[ii], a[iii], product, i, ii, ii)
						count++
					}
				}
			}
		}
	}

	return count
}

/*
const wtf = arr => {
  let count = 0;
  let n = arr.length-1;
  while (n > 2) {
    const doubleCandidates = {};
    const tripleCandidates = {};
    let i = 0;
    while (i <= n) {
      if (tripleCandidates[arr[i]]) {
        count += tripleCandidates[arr[i]];
      }
      for (let j in doubleCandidates) {
        if (j % arr[i] === 0) {
          if (tripleCandidates[j/arr[i]]) {
            tripleCandidates[j/arr[i]] += doubleCandidates[j];
          } else {
            tripleCandidates[j/arr[i]] = doubleCandidates[j];
          }
        }
      }
      if (doubleCandidates[arr[n]/arr[i]]) {
        doubleCandidates[arr[n]/arr[i]] = doubleCandidates[arr[n]/arr[i]]*2;
      } else {
        if (arr[n] % arr[i] === 0) {
          doubleCandidates[arr[n]/arr[i]] = 1;
        }
      }
      i++;
    }
    console.log(doubleCandidates);
    console.log(tripleCandidates);
    n--;
  }
  return count;
}
*/

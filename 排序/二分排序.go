package main

import (
	"fmt"
)

func quickSort(q []int, l, r int) {
	if l >= r {
		return
	}

	i, j := l - 1, r + 1
	mid := q[(i + j) >> 1]
	for i < j {
		for ;; {
			i ++
			if q[i] >= mid {
				break
			}
		}

		for {
			j --
			if q[j] <= mid {
				break
			}
		}

		if i < j {
			q[i], q[j] = q[j], q[i]
		}
	}

	quickSort(q, l, j)
	quickSort(q, j + 1, r)

}

func main() {
	var n int
	fmt.Scanf("%d", &n)

	q := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &q[i])
	}
	quickSort(q, 0, n-1)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", q[i])
	}
}

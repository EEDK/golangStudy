package algorithm

import "fmt"

func Question25(a []int) {
	length := len(a)

	rank := make([]int, length)

	for i := 0; i < length; i++ {
		rank[i] = 1
	}

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if a[j] > a[i] {
				rank[i] = rank[i] + 1
			}
		}
	}

	for i := 0; i < length; i++ {
		fmt.Printf("%d 점 - %d 등 \n", a[i], rank[i])
	}
}

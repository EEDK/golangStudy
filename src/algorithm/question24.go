package algorithm

import (
	"fmt"
)

func Question(a []int) {
	length := len(a)

	num := make(map[int]int, 10)

	for i := 0; i < length; i++ {
		num[a[i]/10] += 1
	}

	fmt.Println(num)

}

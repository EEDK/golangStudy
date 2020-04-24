package algorithm

import (
	"fmt"
)

func Question24(a []int) {
	length := len(a)

	num := make(map[int]int, 10)

	for i := 0; i < length; i++ {
		num[a[i]/10] += 1
	}

	fmt.Println(num)

}

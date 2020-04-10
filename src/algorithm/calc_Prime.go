package algorithm

import "fmt"

func Calc_Prime(n int) {
	num := make(map[int]bool, n+1)
	numMax := int(n ^ (1 / 2))

	for i := 2; i < numMax+1; i++ {
		if num[i] == false {
			for j := i * i; j < n+1; j += i {
				num[j] = true
			}
		}
	}

	for i := 2; i < n+1; i++ {
		if num[i] == false {
			fmt.Print(i, " ")
		}
	}
}

package algorithm

import "fmt"

func Check_prime(n int) {
	for i := 2; i <= n; i++ {
		if Is_n_prime_number(i) {
			fmt.Printf("%d 는 소수\n", i)
		} else {
			fmt.Printf("%d 는 합성수\n", i)
		}
	}
}

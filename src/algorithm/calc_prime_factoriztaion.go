package algorithm

import "fmt"

func Calc_prime_factoriztaion(n int) {
	for i := 2; i <= n; i++ {
		for {
			if n%i == 0 {
				fmt.Printf("%d ", i)
				n = n / i
			} else {
				break
			}
		}
	}
}

package algorithm

import "fmt"

func IsPrime(n int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Printf("약수 : %d\n", i)
		}
	}
}

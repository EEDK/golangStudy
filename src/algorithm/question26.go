package algorithm

import "fmt"

func Question26(n int, k int) {
	d := "0123456789ABCDEFGHIJ"

	if n > 0 {
		Question26(n/k, k)

		fmt.Printf("%c", d[n%k])
	}
}

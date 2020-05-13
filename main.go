package main

import (
	"fmt"
)


func main() {

	var n , s int

	fmt.Scan(&n)
	if n > 50{
		return
	}

	x, err := intScanln(n)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0 ; i < n ; i++{
		if x[i] > 1000000 || x[i] < 0{
			return
		}
	}

	fmt.Scan(&s)
	if s > 1000000 || s < 0{
		return
	}

	checker := 0

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1 ; j++{
			if x[i] < x[i+1] && checker < s{
				temp := x[i+1]
				x[i+1] = x[i]
				x[i] = temp
				checker++
			}
		}
	}

	for i := 0 ; i < n ; i++ {
		fmt.Printf("%d ", x[i])
	}
	fmt.Println();
}

func intScanln(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}
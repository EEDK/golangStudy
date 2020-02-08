package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Sqrt(10))
	fmt.Println(Sqrt(10))
}


func Sqrt(x float64) float64 {
	z := float64(1)

	for i := 0 ; i < 10 ; i++ {
		result := z-(z*z-x)/(2*z)

		if float64(result) == float64(z){
			return result
		}
		z = result
	}
	return z
}

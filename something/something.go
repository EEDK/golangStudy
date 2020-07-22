package something

import "fmt"

func sayHello(){
	fmt.Println("Hello")
}

func SayGoodBye(){
	fmt.Println("GoodBye")
}

func Factorial(n int) int {
	if n == 1 {
		return 1
	} else {
		return Factorial(n - 1) * n
	}
}
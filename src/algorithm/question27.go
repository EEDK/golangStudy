package algorithm

// 반복적
/*func Question27(n int) int {
	fact := 1

	for i := n; i > 1; i-- {
		fact *= i
	}

	return fact
}*/

func Question27(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * Question27(n-1)
	}
}

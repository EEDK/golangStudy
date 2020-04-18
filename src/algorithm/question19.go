package algorithm

func Qustion19(n int, r int) int {
	p := 1
	for i := 1; i <= r; i++ {
		p = p * (n - i + 1) / i
	}

	return p
}

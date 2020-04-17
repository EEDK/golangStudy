package algorithm

func Question18(n int) []int {
	var a []int = make([]int, n)
	a[0] = 1
	a[1] = 1
	for i := 2; i < n; i++ {
		a[i] = a[i/2] + 1
	}
	return a
}

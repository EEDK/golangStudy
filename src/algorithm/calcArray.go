package algorithm

func CalcArray(n int) []int {
	var a []int = make([]int, n)
	a[0] = 1
	a[1] = 1
	a[2] = 1

	for i := 3; i <= n-1; i++ {
		a[i] = a[i-1] + a[i-2] + 1
	}

	return a
}

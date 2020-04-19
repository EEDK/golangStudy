package algorithm

func Qusetion20(n int, r int) int {
	if r == 0 || r == n {
		return 1
	} else {
		return Qusetion20(n-1, r) + Qusetion20(n-1, r-1)
	}
}

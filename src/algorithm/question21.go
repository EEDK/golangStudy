package algorithm

func Qusetion21(n int, r int) int {
	if r == n || r == 0 {
		return 1
	} else {
		return Qusetion20(n, r-1) * (n - r + 1) / r
	}
}

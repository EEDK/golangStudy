package algorithm

func Qustion22(k int, m int) int {
	if k <= 1 {
		return 1 % m
	} else {
		a := Qustion22(k-1, m)
		b := 2 * Qustion22(k-2, m)
		return (a + b) % m
	}
}

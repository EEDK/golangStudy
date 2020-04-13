package algorithm

func Recurssion_Gcd(n int, m int) int {
	if m == 0 {
		return n
	}
	return Recurssion_Gcd(n, n%m)
}

package algorithm

func Gcd(n int, m int) int {
	if n >= m {
	BACK:
		if m > 0 {
			c := m
			m = n % m
			n = c
			goto BACK
		}
		return n
	} else {
	RIGHTBACK:
		if n > 0 {
			c := n
			n = m % n
			m = c
			goto RIGHTBACK
		}
		return m
	}
}

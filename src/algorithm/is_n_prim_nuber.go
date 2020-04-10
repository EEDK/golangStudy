package algorithm

func Is_n_prime_number(n int) bool {

	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
